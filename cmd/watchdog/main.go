package main

import (
	"flag"
	"fmt"

	watchDogServer "github.com/binuud/watchdog/pkg/watchdog"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Define a string flag for the file name
	fileName := flag.String("file", "config.yaml", "Name of the config file (config.yaml) (optional)")
	pVerbose := flag.Bool("v", false, "Explain what's happening while program runs")

	// Parse the flags
	flag.Parse()

	if !*pVerbose {
		// Only log the warning severity or above.
		// if verbose is not set
		log.SetLevel(log.WarnLevel)
	}

	fmt.Println("usage: watchdog --file [filename-with-path]")
	fmt.Println("Using config file ", *fileName)

	// print created using https://www.fancytextpro.com/BigTextGenerator/Cyberlarge
	fmt.Println(`


 _  _  _ _______ _______ _______ _     _ ______   _____   ______
 |  |  | |_____|    |    |       |_____| |     \ |     | |  ____
 |__|__| |     |    |    |_____  |     | |_____/ |_____| |_____|
                                                                
    `)
	fmt.Println("Fetching data... (Single Thread)")
	w := watchDogServer.NewWatchDogService(*fileName)
	w.CheckDomains()
	// w.PrintSummary() // normal print blocks
	w.PrintSummaryTable() // uses 3rd party pretty table

}
