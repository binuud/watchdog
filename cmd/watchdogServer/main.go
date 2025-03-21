package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	protoV1 "github.com/binuud/watchdog/gen/go/v1/watchdog"
	watchDogServer "github.com/binuud/watchdog/pkg/watchdog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	server_grpc_port = flag.Int("grpc_port", 9090, "Watchdog GRPC Serverport, no token required unsecured")
	server_http_port = flag.Int("http_port", 9080, "Watchdog HTTP Server port, no token required unsecured")
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

	logrus.Info("HTTP server started")
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

	protoV1.RegisterWatchDogServer(insecureServer, watchDogServer.NewWatchDogGRPCServer("config.yaml"))

	// Register reflection service on gRPC server.
	reflection.Register(insecureServer)

	logrus.Infof("GRPC server starting...(%d)", *server_grpc_port)
	logrus.Fatalln(insecureServer.Serve(insecureConn))

}

func main() {

	flag.Parse()

	logrus.Info("Giv me one more night !")
	logrus.Info("--------->>>> Starting WatchDog Server Insecure Mode <<<----------")
	logrus.Infof("GRPC Port - (%d), HTTP Port - (%v)", *server_grpc_port, *server_http_port)

	// print created using https://www.fancytextpro.com/BigTextGenerator/Cyberlarge
	logrus.Printf(`
    _  _  _ _______ _______ _______ _     _ ______   _____   ______
    |  |  | |_____|    |    |       |_____| |     \ |     | |  ____
    |__|__| |     |    |    |_____  |     | |_____/ |_____| |_____|
                                                                   
    `)

	go func() {
		runGRPCServer()
	}()

	runHtppServer()

}
