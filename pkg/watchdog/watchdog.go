package watchDogServer

import (
	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/binuud/watchdog/pkg/domainUtils"
	"github.com/sirupsen/logrus"
)

func (s *WatchDogServer) checkDomains() error {
	for _, item := range s.DomainWatch.Domains {
		s.checkDomain(item)
	}
	return nil
}

func (s *WatchDogServer) checkDomain(domain *watchdog.DomainItem) {

	for _, item := range domain.Endpoints {
		responseCode, err := domainUtils.CheckHttpHead(item)
		if err != nil {
			// handle err
		}
		logrus.Printf("Received response %d %s", responseCode, item)
	}

}
