package main

import (
	watchDogServer "github.com/binuud/watchdog/pkg/watchdog"
	"github.com/sirupsen/logrus"
)

func main() {

	// created using https://www.fancytextpro.com/BigTextGenerator/Cyberlarge
	logrus.Printf(`
 _  _  _ _______ _______ _______ _     _ ______   _____   ______
 |  |  | |_____|    |    |       |_____| |     \ |     | |  ____
 |__|__| |     |    |    |_____  |     | |_____/ |_____| |_____|
                                                                
    `)

	w := watchDogServer.NewWatchDogServer("config.yaml")
	w.CheckDomains()
	w.PrintSummary()

}
