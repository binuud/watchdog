package watchDogServer

import (
	"github.com/binuud/watchdog/gen/go/v1/watchdog"
)

type WatchDogServer struct {
	DomainWatch *watchdog.DomainWatch
	Data        []*watchdog.DomainRow
}

func NewWatchDogServer(fileName string) *WatchDogServer {

	serverObj := &WatchDogServer{}
	return serverObj.initFromConfig(fileName)

}
