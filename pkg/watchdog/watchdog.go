package watchDogServer

// Watchdog server
// 1. Watches certificates, connectivity for domains
// 2. Watches git status for each project entry

import (
	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	log "github.com/sirupsen/logrus"
)

type WatchDogService struct {
	RootGitFolder string
	GitHubToken   string

	// cache settings for the watcher
	DomainWatch *watchdog.DomainWatch

	// collection of the domains to be watched and their statuses and summary
	Data []*watchdog.DomainRow

	// collection of git projects to be watched, their statuses, the remote status etc
	GitData []*watchdog.GitProjectRow
}

func NewWatchDogService(fileName string, rootGitFolder string, gitHubToken string) *WatchDogService {

	serverObj := &WatchDogService{}
	w := serverObj.initFromConfig(fileName)
	w.RootGitFolder = rootGitFolder
	w.GitHubToken = gitHubToken
	return w

}

// read contents from the config yaml file
// build the in memory database
// this allows us to add/delete domains once the server is running
func (s *WatchDogService) initFromConfig(fileName string) *WatchDogService {

	domainWatch := &watchdog.DomainWatch{}
	readYaml(fileName, domainWatch)

	// create object to store domain statuses
	domainEntries := make([]*watchdog.DomainRow, 0)

	// create object to store git project statuses
	projectEntries := make([]*watchdog.GitProjectRow, 0)

	for _, item := range domainWatch.Domains {
		domainEntry := s.getDomainEntry(item.Name)
		if domainEntry == nil {
			// domain not in memory
			// create a new entry for same
			domainEntry = &watchdog.DomainRow{
				Domain: item,
				Info:   &watchdog.DomainInfo{},
				Summary: &watchdog.DomainSummary{
					Domain: item,
				},
			}
		} else {
			log.Fatalf("Remove duplicate domain entry %s", item.Name)
		}
		domainEntries = append(domainEntries, domainEntry)
		//		log.Println(item)
	}

	// for each entry of project in config file
	// create a project status entry in memory
	for _, item := range domainWatch.Projects {
		projectEntry := s.getProjectEntry(item.Name)
		if projectEntry == nil {
			// domain not in memory
			// create a new entry for same
			projectEntry = &watchdog.GitProjectRow{
				Project: item,
				Status:  &watchdog.GitProjectStatus{},
			}
		} else {
			log.Fatalf("Remove duplicate project domain entry %s", item.Name)
		}
		projectEntries = append(projectEntries, projectEntry)
	}

	s.DomainWatch = domainWatch
	s.Data = domainEntries
	s.GitData = projectEntries

	return s
}

func (s *WatchDogService) GetByNameOrUUID(name string, uuid string) *watchdog.DomainRow {
	for _, item := range s.Data {
		if item.Domain.Name == name || item.Domain.Uuid == uuid {
			return item
		}
	}
	return nil
}
