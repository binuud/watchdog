package gitUtils

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/google/go-github/v92/github"
	"github.com/sirupsen/logrus"
)

func UpdateGitProjects(projectRow *watchdog.GitProjectRow, rootGitFolder string) error {

	err := SummarizeGitStatus(projectRow, rootGitFolder)

	logrus.Infof("Repo Status %s, %v", projectRow.Project.Name, projectRow.Status)
	logrus.Infof("Project details %v", projectRow)

	return err

}

func GetEmtpyGitProjectStatus() *watchdog.GitProjectStatus {
	return &watchdog.GitProjectStatus{
		BranchName:     "",
		UpstreamBranch: "",
		AheadCount:     0,
		BehindCount:    0,
		IsClean:        true,
		NumModified:    0,
		NumUnstaged:    0,
		NumUntracked:   0,
		NumIgnored:     0,
	}
}

func SummarizeGitStatus(projectRow *watchdog.GitProjectRow, rootGitFolder string) error {

	absolutePath := filepath.Join(rootGitFolder, projectRow.Project.Path)

	projectRow.Status = GetEmtpyGitProjectStatus()

	logrus.Infof("Getting status of git folder %s", absolutePath)

	// get the git repo associated with the path
	repo, err := git.PlainOpen(absolutePath)
	if err != nil {
		return err
	}

	// Get the working tree for the repository
	worktree, err := repo.Worktree()
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

	repoUrl, err := GetRepoMetadata(repo)
	if err != nil {
		logrus.Errorf("Cannot get git repo meta details for repo %s, err %v", absolutePath, err)
		return err
	}

	projectRow.Project.Remoteurl = repoUrl

	// read the stash count
	stashCount, err := GetStashCount(worktree)
	if err != nil {
		return err
	}
	projectRow.Status.NumStashes = int64(stashCount)

	status, err := GetGitStatus(worktree)
	if err != nil {
		logrus.Errorf("Cannot get git status for repo %s, err %v", absolutePath, err)
		return err
	}

	if status == nil {
		return nil
	}

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

		fmt.Printf("File: %s | Staging: %c | Worktree: %c\n",
			filePath, fileStatus.Staging, fileStatus.Worktree)

	}

	return nil

}

func GetGitStatus(worktree *git.Worktree) (git.Status, error) {

	// Retrieve the current status of the working tree
	status, err := worktree.Status()
	if err != nil {
		return nil, err
	}

	// Check if the working tree is clean
	if status.IsClean() {
		fmt.Println("Working tree is clean. Nothing to commit.")
		return nil, nil
	}

	return status, nil

}

func GetStashCount(worktree *git.Worktree) (int, error) {

	// get stash count
	// Locate the reflog file for stashes
	// path: .git/logs/refs/stash
	stashLogPath := filepath.Join(worktree.Filesystem().Root(), ".git", "logs", "refs", "stash")
	// Open the reflog file
	file, err := os.Open(stashLogPath)
	if os.IsNotExist(err) {
		return 0, nil // No stashes exist yet
	}
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Count the number of lines (each line is one stash)
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	return count, scanner.Err()
}

func GetGitUser(repo *git.Repository) (string, error) {

	// Fetch the repository-specific (local) configuration
	cfg, err := repo.Config()
	logrus.Infof("Config %v", cfg.User)
	if err != nil {
		return "", err
	}

	return cfg.User.Name, nil

}

// getRepoMetadata inspects the local .git config to find the owner and repo name
func GetRepoMetadata(localRepo *git.Repository) (string, error) {

	// Fetch the 'origin' remote configuration
	remote, err := localRepo.Remote("origin")
	if err != nil {
		return "", fmt.Errorf("could not find 'origin' remote: %w", err)
	}

	// 3. Read the remote URL (supports HTTPS and SSH formats)
	if len(remote.Config().URLs) == 0 {
		return "", fmt.Errorf("no remote URLs configured for origin")
	}
	url := remote.Config().URLs[0]

	return url, nil
}

func GetIssuesPullrequests(path string, owner string, repo string) {

	ctx := context.Background()

	// Automatically extract repository information from your local environment
	// owner, repo, err := GetRepoMetadata(path)
	// if err != nil {
	// 	log.Fatalf("Error reading current project context: %v", err)
	// }
	// fmt.Printf("Detected Local Project Context: %s/%s\n", owner, repo)

	// Initialize a standard GitHub client.
	// For production, authenticate your client to prevent rate limits: github.NewClient(nil).WithAuthToken("your_token")
	client, err := github.NewClient()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// owner := "go-git"
	// repo := "go-git"

	// 1. Fetch pure Open Pull Requests count using a search query
	prQuery := fmt.Sprintf("repo:%s/%s is:pr state:open", owner, repo)
	prResult, _, err := client.Search.Issues(ctx, prQuery, &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 1}, // We only care about the Total count metadata
	})
	if err != nil {
		log.Fatalf("Error searching pull requests: %v", err)
	}
	openPRs := prResult.GetTotal()

	// 2. Fetch pure Open Issues count using a search query
	issueQuery := fmt.Sprintf("repo:%s/%s is:issue state:open", owner, repo)
	issueResult, _, err := client.Search.Issues(ctx, issueQuery, &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	})
	if err != nil {
		log.Fatalf("Error searching issues: %v", err)
	}
	openIssues := issueResult.GetTotal()

	// Output the separated metrics
	fmt.Printf("Repository: %s/%s\n", owner, repo)
	fmt.Printf("Pure Open Issues: %d\n", openIssues)
	fmt.Printf("Open Pull Requests: %d\n", openPRs)

}

func check_remote_status() {

	// 1. Open the local repository
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Failed to open repo: %v", err)
	}

	// 2. Fetch updates from the remote 'origin'
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
	})
	// Ignore errors if the repository is already up to date
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		log.Fatalf("Failed to fetch from remote: %v", err)
	}

	// 3. Get the current local HEAD reference
	headRef, err := repo.Head()
	if err != nil {
		log.Fatalf("Failed to get HEAD: %v", err)
	}

	// 4. Resolve the remote-tracking reference (assuming 'origin/master')
	remoteRefName := plumbing.ReferenceName("refs/remotes/origin/master")
	remoteRef, err := repo.Reference(remoteRefName, true)
	if err != nil {
		log.Fatalf("Failed to get remote reference: %v", err)
	}

	// 5. Compare local and remote hashes
	if headRef.Hash() == remoteRef.Hash() {
		fmt.Println("Your local branch is up to date with the remote.")
	} else {
		fmt.Printf("Changes found! Local: %s, Remote: %s. Pull recommended.\n",
			headRef.Hash().String()[:7], remoteRef.Hash().String()[:7])
	}

}
