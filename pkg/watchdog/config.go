package watchDogServer

import (
	"io"
	"os"

	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func readYaml(fileName string, domainWatch *watchdog.DomainWatch) {

	logrus.Infof("Reading config file %s", fileName)
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		logrus.Fatalf("Cannot read %s, %v", fileName, err)
	}

	err = yaml.Unmarshal(yamlFile, domainWatch)
	if err != nil {
		logrus.Fatalf("Malformed yaml file %s, %v", fileName, err)
	}
}

func writeYaml(domainWatch *watchdog.DomainWatch, fileName string) {

	yamlData, err := yaml.Marshal(domainWatch)
	if err != nil {
		logrus.Fatalf("Cannot marshal domains into yaml %v", err)
	}

	f, err := os.Create(fileName)
	if err != nil {
		logrus.Fatalf("Cannot Create %s, %v", fileName, err)
	}
	defer f.Close()

	_, err = io.Writer.Write(f, yamlData)
	if err != nil {
		logrus.Fatalf("Cannot write %s, %v", fileName, err)
	}

}
