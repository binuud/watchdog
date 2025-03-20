package watchDogServer

import (
	"github.com/binuud/watchdog/gen/go/v1/watchdog"
)

type WatchDogServer struct {
	DomainWatch *watchdog.DomainWatch
}

func NewWatchDogServer(fileName string) *WatchDogServer {
	domainWatch := &watchdog.DomainWatch{}
	readYaml(fileName, domainWatch)
	return &WatchDogServer{
		DomainWatch: domainWatch,
	}
}
