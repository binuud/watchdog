package main

import (
	"flag"
	"fmt"

	watchDogServer "github.com/binuud/watchdog/pkg/watchdog"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	// Define a string flag for the file name
	fileName := flag.String("file", "config.yaml", "Name of the config file (config.yaml) (optional)")

	// Parse the flags
	flag.Parse()

	fmt.Println("usage: watchdog --file [filename-with-path]")
	fmt.Println("Using config file ", *fileName)

	// print created using https://www.fancytextpro.com/BigTextGenerator/Cyberlarge
	fmt.Println(`


 _  _  _ _______ _______ _______ _     _ ______   _____   ______
 |  |  | |_____|    |    |       |_____| |     \ |     | |  ____
 |__|__| |     |    |    |_____  |     | |_____/ |_____| |_____|
                                                                
    `)

	w := watchDogServer.NewWatchDogService(*fileName)
	w.CheckDomains()
	// w.PrintSummary() // normal print blocks
	w.PrintSummaryTable() // uses 3rd party pretty table

}
