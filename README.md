# GoGitAnalyst
GoGitAnalyst
GoGitAnalyst is a command-line tool to monitor GitHub repository statistics. It collects information about the repository's participants, branches, commits, and displays the statistics in a readable format.

#Feature
- Fetches repository participants and their roles.
- Lists branches and their first commit timestamps.
- Displays commit details.
- Collects user commit statistics.

#Installation
Before installing, make sure you have the following tools installed:
- Go (version 1.18+): You need to have Go installed to build the project. You can install it from here.
- GitHub OAuth Token: You need a GitHub OAuth token to access the repository data.

1 - Clone the repository
```bash
git clone https://github.com/yourusername/GoGitAnalyst.git
cd GoGitAnalyst
```

2 - Install dependencies
GoGitAnalyst uses external packages, so you will need to install them using Go's package manager.
Run the following command to download the required dependencies:

```bash
go mod init GoGitAnalyst
go mod tidy
```
#Usage
1 - Generate OAuth Token
To access private GitHub repositories or make authenticated API requests, you need to generate a GitHub OAuth token.
- Go to GitHub's Personal Access Tokens page.
- Click on "Generate new token".
- Select the necessary scopes (for example, repo to access private repositories).
- Generate the token and copy it. (BE SURE TO SAVE IT SOMEWHERE, AS IT IS DISPLAYED ONLY ONCE RIGHT WHEN YOU GENERATE IT!)

2 - Run the Tool
Once you have your GitHub OAuth token, you can use the CLI tool to gather repository stats.

```bash
go run main.go --token <YOUR_OAUTH_TOKEN> --owner <REPO_OWNER> --repo <REPO_NAME>
```

#Flags
--token, -t: GitHub OAuth Token (required).
--owner, -o: GitHub repository owner (required).
--repo, -r: GitHub repository name (required).

#Example:

```bash
go run main.go --token abc123 --owner mygithubuser --repo myrepository
Example Output
The output of the command will show:
Users and Roles: Lists all participants with their roles (Owner, Collaborator).
Branches: Displays the branches and the timestamp of the first commit.
Commits: Lists commits with their timestamps, hashes, commit messages, and the name of the committer.
User Stats: Shows commit statistics per user.
```

#Example:

```bash
Users and Roles:
- mygithubuser | Owner
- collaborator1 | Collaborator
- collaborator2 | Collaborator

Branches:
- main | Mon, 12 Apr 2021 12:00:00 UTC
- feature-branch | Wed, 14 Apr 2021 15:30:00 UTC

Commits:
Mon, 12 Apr 2021 12:00:00 UTC | abc123 | mygithubuser | Initial commit
Wed, 14 Apr 2021 15:30:00 UTC | def456 | collaborator1 | Added new feature

User Stats:
- mygithubuser | Owner | 1 commits
- collaborator1 | Collaborator | 1 commits
```

#Troubleshooting
- If you encounter issues fetching data, ensure your OAuth token has the correct permissions for the repository.
- If the program cannot find the repository or encounters an error, double-check the repository owner and name.
