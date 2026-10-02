package watchDogServer

import (
	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/binuud/watchdog/pkg/gitUtils"
	log "github.com/sirupsen/logrus"
)

func (s *WatchDogService) CheckProjects() error {

	log.Infof("Fetching details of all git projects")
	for _, projectEntry := range s.GitData {

		// get domain details
		// persist data in domain Entry for summarization and storage
		s.getProjectDetails(projectEntry)

	}

	return nil

}

// get the git project row corresponding to the given git project name
// from the data object
func (s *WatchDogService) getProjectEntry(projectName string) *watchdog.GitProjectRow {
	for _, item := range s.GitData {
		if projectName == item.Project.Name {
			return item
		}
	}
	return nil
}

// get domain details like certificates
// ip resolv and reachablilty of a domain
func (s *WatchDogService) getProjectDetails(projectEntry *watchdog.GitProjectRow) {
	log.Infof("Project details for Name: %s, Owner: %s", projectEntry.Project.Name, projectEntry.Project.User)
	gitUtils.UpdateGitProjects(projectEntry, s.RootGitFolder)
}
