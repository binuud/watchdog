package watchDogServer

import (
	"testing"

	"github.com/binuud/watchdog/gen/go/v1/watchdog"
)

func TestWriteYaml(t *testing.T) {

	testConfigFile := "testSampleConfig.yaml"

	domainWatch := &watchdog.DomainWatch{
		Name:            "TestConfig",
		RefreshInterval: 86400,
	}
	domains := make([]*watchdog.DomainItem, 0)

	domainItem := &watchdog.DomainItem{
		Name: "www.google.com",
	}
	domains = append(domains, domainItem)

	domainItem = &watchdog.DomainItem{
		Name: "www.gmail.com",
	}
	domains = append(domains, domainItem)
	domainWatch.Domains = domains

	writeYaml(domainWatch, testConfigFile)

	t.Log("No panics seen during write of yaml config")

}

func TestReadYaml(t *testing.T) {

	testConfigFile := "testSampleConfig.yaml"
	domainWatch := &watchdog.DomainWatch{}
	readYaml(testConfigFile, domainWatch)

	t.Log("No panics seen during read of yaml config")
}
