package gitUtils

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/client"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
	"github.com/google/go-github/v92/github"
	"github.com/sirupsen/logrus"
)

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
		CommonAncestor: "",
	}
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

func GetGitOpenIssues(client *github.Client, owner *string, repo *string) (int, error) {

	ctx := context.Background()

	issueQuery := fmt.Sprintf("repo:%s/%s is:issue state:open", *owner, *repo)
	issueResult, _, err := client.Search.Issues(ctx, issueQuery, &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	})
	if err != nil {
		return -1, err
	}

	return issueResult.GetTotal(), nil

}

func GetGitOpenPR(client *github.Client, owner *string, repo *string) (int, error) {

	ctx := context.Background()
	prQuery := fmt.Sprintf("repo:%s/%s is:pr state:open", *owner, *repo)
	prResult, _, err := client.Search.Issues(ctx, prQuery, &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 1}, // We only care about the Total count metadata
	})
	if err != nil {
		return -1, err
	}

	return prResult.GetTotal(), nil

}

func CheckGitRemoteStatus(repo *git.Repository, githubToken string) (commonAncestor string, behind int, ahead int, err error) {

	commonAncestor = ""
	behind = 0
	ahead = 0

	// Fetch updates from the remote 'origin'
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		ClientOptions: []client.Option{
			// Explicitly use TokenAuth to bypass the password authentication block
			client.WithHTTPAuth(&http.BasicAuth{
				Username: "x-access-token", // Best practice for GitHub API/Actions tokens
				Password: githubToken,      // The actual token string
			}),
		},
	})
	// Ignore errors if the repository is already up to date
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		logrus.Errorf("CheckGitRemoteStatus - Failed to fetch from remote: %v", err)
		return
	} // else if errors.Is(err, git.NoErrAlreadyUpToDate) {
	// 	return
	// }

	// Get the current local HEAD reference
	headRef, err := repo.Head()
	logrus.Infof("CheckGitRemoteStatus heardre %s", headRef.Name())
	if err != nil {
		logrus.Errorf("CheckGitRemoteStatus - Failed to get HEAD: %v", err)
		return
	}

	// 4. Resolve the remote-tracking reference (assuming 'origin/master')
	remoteRefName := plumbing.ReferenceName("refs/remotes/origin/master")
	remoteRef, err := repo.Reference(remoteRefName, true)
	if err != nil {
		logrus.Errorf("Failed to get remote reference: %v", err)
		return
	}

	// 5. Compare local and remote hashes
	if headRef.Hash() == remoteRef.Hash() {
		logrus.Println("Your local branch is up to date with the remote.")
	} else {
		logrus.Printf("Changes found! Local: %s, Remote: %s. Pull recommended.\n",
			headRef.Hash().String()[:7], remoteRef.Hash().String()[:7])
		commonAncestor, behind, ahead, err = countAheadBehind(repo, headRef.Hash(), remoteRef.Hash())
		if err != nil {
			return
		}
	}

	return
}

// Helper to count commits ahead and behind via commit history walking
func countAheadBehind(repo *git.Repository, localHash, remoteHash plumbing.Hash) (commonAncestor string, behind int, ahead int, err error) {

	commonAncestor = ""
	behind = 0
	ahead = 0

	localCommit, err := repo.CommitObject(localHash)
	if err != nil {
		logrus.Errorf("countAheadBehind: cannot read localhash commit %v", err)
		return
	}

	remoteCommit, err := repo.CommitObject(remoteHash)
	if err != nil {
		logrus.Errorf("countAheadBehind: cannot read remoteHash commit %v", err)
		return
	}

	// Find merge base (common ancestor)
	mergeBases, err := localCommit.MergeBase(remoteCommit)
	if err != nil || len(mergeBases) == 0 {
		logrus.Errorf("countAheadBehind: no common ancestor %v", err)
		return
	}
	baseHash := mergeBases[0].Hash
	commonAncestor = baseHash.String()[:7]

	// Count commits from local to merge base (Ahead)
	ahead, err = countCommitsBetween(localCommit, baseHash)
	if err != nil {
		logrus.Errorf("countAheadBehind: cannot countCommitsBetween localhash %v", err)
		return
	}

	// Count commits from remote to merge base (Behind)
	behind, err = countCommitsBetween(remoteCommit, baseHash)
	if err != nil {
		logrus.Errorf("countAheadBehind: cannot countCommitsBetween remotehash %v", err)
		return
	}

	return
}

// Walk commit tree from start down to stopHash
func countCommitsBetween(start *object.Commit, stopHash plumbing.Hash) (int, error) {
	count := 0
	// create an interator
	seen := make(map[plumbing.Hash]bool)

	// NewCommitPreorderIter is the correct constructor for history walking
	walker := object.NewCommitPreorderIter(start, seen, nil)

	err := walker.ForEach(func(c *object.Commit) error {
		if c.Hash == stopHash {
			// Stop counting when we hit the common ancestor
			return storer.ErrStop
		}
		count++
		return nil
	})

	return count, err
}
