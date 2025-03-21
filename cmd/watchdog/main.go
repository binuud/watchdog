package main

import (
	"flag"
	"fmt"

	watchDogServer "github.com/binuud/watchdog/pkg/watchdog"
	"github.com/sirupsen/logrus"
)

func main() {

	// Define a string flag for the file name
	fileName := flag.String("file", "config.yaml", "Name of the config file (config.yaml) (optional)")

	// Parse the flags
	flag.Parse()

	fmt.Println("usage: watchdog --file [filename-with-path]")
	fmt.Printf("\n Using config file %s", *fileName)

	// print created using https://www.fancytextpro.com/BigTextGenerator/Cyberlarge
	logrus.Printf(`
 _  _  _ _______ _______ _______ _     _ ______   _____   ______
 |  |  | |_____|    |    |       |_____| |     \ |     | |  ____
 |__|__| |     |    |    |_____  |     | |_____/ |_____| |_____|
                                                                
    `)

	w := watchDogServer.NewWatchDogService(*fileName)
	w.CheckDomains()
	w.PrintSummary()

}
