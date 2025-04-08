package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v43/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var token string
var repoOwner string
var repoName string

// getGitHubClient creates a GitHub client with an OAuth token.
func getGitHubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	return client
}

// getRepoStats collects statistics for the repository: participants, branches, and commits.
func getRepoStats(client *github.Client, owner, repo string) {
	ctx := context.Background()

	// Get participants
	participants, _, err := client.Repositories.ListCollaborators(ctx, owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a map for users and their statistics
	userStats := make(map[string]struct {
		Role        string
		CommitCount int
	})

	// Get repository owner information
	repoDetails, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	// Get repository owner login
	repoOwnerLogin := *repoDetails.Owner.Login

	// Get user roles and their information
	fmt.Println("Users and Roles:")
	for _, user := range participants {
		// Exclude the "GitHub" user
		if *user.Login == "GitHub" {
			continue
		}

		// Assign role "Owner" if the user is the repository owner
		role := "Collaborator"
		if *user.Login == repoOwnerLogin {
			role = "Owner"
		}

		userStats[*user.Login] = struct {
			Role        string
			CommitCount int
		}{Role: role}
		fmt.Printf("- %s | %s\n", *user.Login, role)
	}

	// Get branches
	branches, _, err := client.Repositories.ListBranches(ctx, owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print information about branches
	fmt.Println("\nBranches:")
	for _, branch := range branches {
		branchName := *branch.Name
		// Get the first commit on the branch
		commits, _, err := client.Repositories.ListCommits(ctx, owner, repo, &github.CommitsListOptions{SHA: branchName})
		if err != nil {
			log.Fatal(err)
		}
		if len(commits) > 0 {
			firstCommitTime := commits[0].Commit.Committer.Date.Format(time.RFC1123)
			fmt.Printf("- %s | %s\n", branchName, firstCommitTime)
		} else {
			fmt.Printf("- %s | No commits found\n", branchName)
		}
	}

	// Get commits
	commits, _, err := client.Repositories.ListCommits(ctx, owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print information about commits and update user statistics
	fmt.Println("\nCommits:")
	for _, commit := range commits {
		commitTime := commit.Commit.Committer.Date.Format(time.RFC1123) // Format commit time
		commitHash := *commit.SHA
		commitMessage := *commit.Commit.Message
		committerName := *commit.Commit.Committer.Name

		// Update user statistics
		if stats, exists := userStats[committerName]; exists {
			stats.CommitCount++
			userStats[committerName] = stats
		} else {
			userStats[committerName] = struct {
				Role        string
				CommitCount int
			}{Role: "Contributor", CommitCount: 1} // If the user is not found, consider it as their first commit
		}

		// If the commit message is empty, display only the hash
		if commitMessage == "" {
			commitMessage = "No message"
		}

		// Print commit information
		fmt.Printf("%s | %s | %s | %s\n", commitTime, commitHash, committerName, commitMessage)
	}

	// Print user statistics
	fmt.Println("\nUser Stats:")
	for user, stats := range userStats {
		// Exclude "GitHub"
		if user == "GitHub" {
			continue
		}
		fmt.Printf("- %s | %s | %d commits\n", user, stats.Role, stats.CommitCount)
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "monitoring-cli",
		Short: "A tool to monitor GitHub repository stats",
		Run: func(cmd *cobra.Command, args []string) {
			// Create GitHub client
			client := getGitHubClient(token)

			// Get repository stats
			getRepoStats(client, repoOwner, repoName)
		},
	}

	// Add flags for CLI input
	rootCmd.Flags().StringVarP(&token, "token", "t", "", "GitHub OAuth Token")
	rootCmd.Flags().StringVarP(&repoOwner, "owner", "o", "", "Repository owner (e.g., 'MyName')")
	rootCmd.Flags().StringVarP(&repoName, "repo", "r", "", "Repository name (e.g., 'MyRepoIsCool')")

	// Mark flags as required
	rootCmd.MarkFlagRequired("token")
	rootCmd.MarkFlagRequired("owner")
	rootCmd.MarkFlagRequired("repo")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
