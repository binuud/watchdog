package watchDogServer

import (
	"time"

	"github.com/sirupsen/logrus"
)

const MIN_REFRESH_INTERVAL = 300 // 5 minutes or 300 seconds

func (s *WatchDogService) Watcher() {
	logrus.Printf("Domain watch %v", s.DomainWatch)
	if s.DomainWatch.RefreshInterval < MIN_REFRESH_INTERVAL {
		logrus.Warnf("Refresh Interval cannot be less than 60 seconds, setting to 60")
		s.DomainWatch.RefreshInterval = MIN_REFRESH_INTERVAL
	}
	// run checkDomain once
	s.CheckDomains()

	ticker := time.NewTicker(time.Duration(s.DomainWatch.RefreshInterval) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		logrus.Infof("Collecting domain data")
		s.CheckDomains()
	}

}
