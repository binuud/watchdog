package watchDogServer

import (
	"path/filepath"

	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/binuud/watchdog/pkg/gitUtils"
	"github.com/go-git/go-git/v6"
	"github.com/google/go-github/v92/github"
	"github.com/sirupsen/logrus"
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
	s.UpdateGitProjects(projectEntry)
}

func (s *WatchDogService) UpdateGitProjects(projectRow *watchdog.GitProjectRow) error {

	err := s.SummarizeGitStatus(projectRow)

	logrus.Infof("Repo Status %s, %v", projectRow.Project.Name, projectRow.Status)
	logrus.Infof("Project details %v", projectRow)

	return err

}

func (s *WatchDogService) SummarizeGitStatus(projectRow *watchdog.GitProjectRow) error {

	absolutePath := filepath.Join(s.RootGitFolder, projectRow.Project.Path)

	projectRow.Status = gitUtils.GetEmtpyGitProjectStatus()

	logrus.Infof("Getting status of git folder %s", absolutePath)

	// get the git repo associated with the path
	localRepo, err := git.PlainOpen(absolutePath)
	if err != nil {
		return err
	}

	// Get the working tree for the repository
	worktree, err := localRepo.Worktree()
	if err != nil {
		return err
	}

	// Since many users use a global config for user, do not read from local repo config

	// user, err := GetGitUser(repo)
	// if err != nil {
	// 	logrus.Errorf("Cannot get git user details for repo %s, err %v", absolutePath, err)
	// 	return err
	// }

	// projectRow.Project.User = user

	repoUrl, err := gitUtils.GetRepoMetadata(localRepo)
	if err != nil {
		logrus.Errorf("Cannot get git repo meta details for repo %s, err %v", absolutePath, err)
		return err
	}

	projectRow.Project.Remoteurl = repoUrl

	// read the stash count
	stashCount, err := gitUtils.GetStashCount(worktree)
	if err != nil {
		// return err
	}
	projectRow.Status.NumStashes = int64(stashCount)

	status, err := gitUtils.GetGitStatus(worktree)
	if err != nil {
		logrus.Errorf("Cannot get git status for repo %s, err %v", absolutePath, err)
		// return err
	}

	if status == nil {
		// return nil
	} else {
		projectRow.Status.IsClean = false

		for filePath, fileStatus := range status {

			// fileStatus.Staging tells you the status in the index
			// fileStatus.Worktree tells you the status in the working directory

			switch fileStatus.Worktree {
			case git.Added:
				projectRow.Status.NumUntracked++
			case git.Untracked:
				projectRow.Status.NumUntracked++
			case git.Modified:
				projectRow.Status.NumModified++
			}

			logrus.Printf("File: %s | Staging: %c | Worktree: %c\n",
				filePath, fileStatus.Staging, fileStatus.Worktree)

		}
	}

	commonAncestor, behind, ahead, err := gitUtils.CheckGitRemoteStatus(localRepo, s.GitHubToken)
	if err != nil {
		logrus.Errorf("Error reading git remote status: %v", err)
	}

	projectRow.Status.CommonAncestor = commonAncestor
	projectRow.Status.BehindCount = int32(behind)
	projectRow.Status.AheadCount = int32(ahead)

	gitClient, err := github.NewClient(github.WithAuthToken(s.GitHubToken))
	logrus.Infof("Using github token %s", s.GitHubToken)
	if err != nil {
		logrus.Errorf("Error creating client: %v", err)
	} else {
		numIssues, err := gitUtils.GetGitOpenIssues(gitClient, &projectRow.Project.User, &projectRow.Project.Name)
		if err != nil {
			projectRow.Status.NumIssuesOpen = -1
			logrus.Errorf("Error reading open issues: %v", err)
		} else {
			projectRow.Status.NumIssuesOpen = int64(numIssues)
		}

		numPR, err := gitUtils.GetGitOpenPR(gitClient, &projectRow.Project.User, &projectRow.Project.Name)
		if err != nil {
			projectRow.Status.NumPrOpen = -1
			logrus.Errorf("Error reading open PR: %v", err)
		} else {
			projectRow.Status.NumPrOpen = int64(numPR)
		}
	}

	return nil

}
