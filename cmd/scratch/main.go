package main

import (
	"context"
	"fmt"
	"log"
	"errors"
	"regexp"


	"github.com/go-git/go-git/v6" // Use the latest v6 module version
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/google/go-github/v92/github"	

)

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


func check_git_status() {
	// 1. Open the local repository in the current directory
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Failed to open repository: %v", err)
	}

	// 2. Get the working tree for the repository
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Failed to get worktree: %v", err)
	}

	// 3. Retrieve the current status of the working tree
	status, err := worktree.Status()
	if err != nil {
		log.Fatalf("Failed to get status: %v", err)
	}

	// 4. Check if the working tree is clean
	if status.IsClean() {
		fmt.Println("Working tree is clean. Nothing to commit.")
		return
	}

	// 5. Print out the changed files and their specific statuses
	fmt.Println("Changed files:")
	for filePath, fileStatus := range status {
		// fileStatus.Staging tells you the status in the index
		// fileStatus.Worktree tells you the status in the working directory
		fmt.Printf("File: %s | Staging: %c | Worktree: %c\n", 
			filePath, fileStatus.Staging, fileStatus.Worktree)
	}
}

// getRepoMetadata inspects the local .git config to find the owner and repo name
func getRepoMetadata() (owner string, repo string, err error) {
	// 1. Open the current directory's git repository
	localRepo, err := git.PlainOpen(".")
	if err != nil {
		return "", "", fmt.Errorf("failed to open local git repo: %w", err)
	}

	// 2. Fetch the 'origin' remote configuration
	remote, err := localRepo.Remote("origin")
	if err != nil {
		return "", "", fmt.Errorf("could not find 'origin' remote: %w", err)
	}

	// 3. Read the remote URL (supports HTTPS and SSH formats)
	if len(remote.Config().URLs) == 0 {
		return "", "", fmt.Errorf("no remote URLs configured for origin")
	}
	url := remote.Config().URLs[0]

	// Regex to extract owner and repo from:
	// - https://github.com
	// - git@github.com:owner/repo.git
	re := regexp.MustCompile(`github\.com[:/](.+?)/(.+?)(\.git)?$`)
	matches := re.FindStringSubmatch(url)
	if len(matches) < 3 {
		return "", "", fmt.Errorf("unable to parse GitHub owner/repo from URL: %s", url)
	}

	return matches[1], matches[2], nil
}


func get_issues_pullrequests() {

	ctx := context.Background()
	
	// Automatically extract repository information from your local environment
	owner, repo, err := getRepoMetadata()
	if err != nil {
		log.Fatalf("Error reading current project context: %v", err)
	}
	fmt.Printf("Detected Local Project Context: %s/%s\n", owner, repo)


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

func main() {
	check_git_status()
	check_remote_status()
	get_issues_pullrequests()
}