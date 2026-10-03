package main

import (
	"flag"
	"fmt"

	watchDogServer "github.com/binuud/watchdog/pkg/watchdog"
	log "github.com/sirupsen/logrus"
)

var (
	config_filename = flag.String("config_file", "./configs/config.yaml", "Config file path (config.yaml) (optional)")
	git_root_folder = flag.String("code_root_folder", "/app", "Root folder where all git projects reside (optional)")
	GITHUB_TOKEN    = flag.String("GITHUB_TOKEN", "", "GitHub Token, needed for api access to remote github (optional)")
	pVerbose        = flag.Bool("v", false, "Detailed logs")
)

func main() {

	fmt.Println("Usage: watchdogServer -h   For Help")
	fmt.Print("\n\n")

	// Parse the flags
	flag.Parse()

	if !*pVerbose {
		// Only log the warning severity or above.
		// if verbose is not set
		log.SetLevel(log.WarnLevel)
	}

	fmt.Println("Using config file ", *config_filename)

	// print created using https://www.fancytextpro.com/BigTextGenerator/Cyberlarge
	fmt.Println(`


 _  _  _ _______ _______ _______ _     _ ______   _____   ______
 |  |  | |_____|    |    |       |_____| |     \ |     | |  ____
 |__|__| |     |    |    |_____  |     | |_____/ |_____| |_____|
                                                                
    `)
	fmt.Println("Fetching data... (Single Thread)")
	w := watchDogServer.NewWatchDogService(*config_filename, *git_root_folder, *GITHUB_TOKEN)
	w.CheckDomains()
	// w.PrintSummary() // normal print blocks
	w.PrintSummaryTable() // uses 3rd party pretty table

}
