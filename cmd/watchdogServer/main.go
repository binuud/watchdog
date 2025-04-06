package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"

	protoV1 "github.com/binuud/watchdog/gen/go/v1/watchdog"
	watchDogServer "github.com/binuud/watchdog/pkg/watchdog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	server_grpc_port = flag.Int("grpc_port", 9090, "Watchdog GRPC Serverport, no token required unsecured")
	server_http_port = flag.Int("http_port", 9080, "Watchdog HTTP Server port, no token required unsecured")
	config_filename  = flag.String("file", "./configs/config.yaml", "Config file path (config.yaml) (optional)")
)

// starting the http server from the GPRC generated code
func runHtppServer() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		fmt.Sprintf("0.0.0.0:%d", *server_grpc_port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwMux := runtime.NewServeMux()

	err = protoV1.RegisterWatchDogHandler(ctx, gwMux, conn)
	if err != nil {
		log.Fatalln(err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *server_http_port),
		Handler: gwMux,
	}

	log.Info("HTTP server started")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Fatalln(gwServer.ListenAndServe())

}

// start the grpc server
func runGRPCServer() {

	insecureConn, err := net.Listen("tcp", fmt.Sprintf(":%d", *server_grpc_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	insecureServer := grpc.NewServer()

	protoV1.RegisterWatchDogServer(insecureServer, watchDogServer.NewWatchDogGRPCServer(*config_filename))

	// Register reflection service on gRPC server.
	reflection.Register(insecureServer)

	log.Infof("GRPC server starting...(%d)", *server_grpc_port)
	log.Fatalln(insecureServer.Serve(insecureConn))

}

func main() {

	fmt.Println("\nusage: watchdogServer -h")
	fmt.Println(" -h   For Help")
	fmt.Print("\n\n")
	pVerbose := flag.Bool("v", false, "Detailed logs")

	// Parse the flags
	flag.Parse()

	if !*pVerbose {
		// Only log the warning severity or above.
		// if verbose is not set
		log.SetLevel(log.WarnLevel)
	}

	fmt.Println("Using config file ", *config_filename)

	log.Info("Give me one more night !")
	log.Info("--------->>>> Starting WatchDog Server Insecure Mode <<<----------")
	log.Infof("GRPC Port - (%d), HTTP Port - (%v)", *server_grpc_port, *server_http_port)

	// print created using https://www.fancytextpro.com/BigTextGenerator/Cyberlarge
	fmt.Println(`
    _  _  _ _______ _______ _______ _     _ ______   _____   ______
    |  |  | |_____|    |    |       |_____| |     \ |     | |  ____
    |__|__| |     |    |    |_____  |     | |_____/ |_____| |_____|
                                                                   
    `)

	go func() {
		runGRPCServer()
	}()

	runHtppServer()

}
