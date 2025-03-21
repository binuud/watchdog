package watchDogServer

import (
	"testing"

	"github.com/binuud/watchdog/gen/go/v1/watchdog"
)

func createDomainWatch() *watchdog.DomainWatch {

	domainWatch := &watchdog.DomainWatch{
		Name:            "TestConfig",
		RefreshInterval: 86400,
	}
	domains := make([]*watchdog.DomainItem, 0)

	domainItem := &watchdog.DomainItem{
		Name:      "www.google.com",
		Endpoints: []string{"https://www.google.com"},
	}
	domains = append(domains, domainItem)

	domainItem = &watchdog.DomainItem{
		Name:      "www.gmail.com",
		Endpoints: []string{"https://www.gmail.com"},
	}
	domains = append(domains, domainItem)
	domainWatch.Domains = domains
	return domainWatch

}

func TestCheckDomains(t *testing.T) {

	testConfigFile := "testSampleConfig.yaml"
	domainWatch := createDomainWatch()
	writeYaml(domainWatch, testConfigFile)

	s := NewWatchDogService(testConfigFile)
	err := s.CheckDomains()
	if err != nil {
		t.Errorf("Error in getting domain information (%v)", err)
	}
	t.Logf("Success in checkDomains")

}
