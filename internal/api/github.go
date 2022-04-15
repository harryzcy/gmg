package api

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cli/oauth"
	"github.com/google/go-github/v43/github"
	"github.com/harryzcy/gmg/internal/storage"
	"golang.org/x/oauth2"
)

// Login obtains OAuth tokens from GitHub
func Login() error {
	storage.InitDefault()

	token := storage.GetToken(storage.TokenKindCLI)
	if token != nil {
		fmt.Println("You already have a CLI token, skipping...")
	} else {
		fmt.Println("Generating GitHub OAuth token for CLI...")
		flow := &oauth.Flow{
			Host:     oauth.GitHubHost("https://github.com"),
			ClientID: os.Getenv("GITHUB_CLIENT_ID"),
			Scopes:   []string{"repo"},
		}

		accessToken, err := flow.DetectFlow()
		if err != nil {
			return err
		}

		err = storage.StoreToken(storage.TokenKindCLI, accessToken)
		if err != nil {
			return err
		}
		fmt.Println("Successfully authenticated with GitHub")
	}

	token = storage.GetToken(storage.TokenKindGitHubGitea)
	if token != nil {
		fmt.Println("You already have a GitHub OAuth token for Gitea, skipping...")
	} else {
		fmt.Println("Generating GitHub OAuth token for Gitea...")
		flow := &oauth.Flow{
			Host:     oauth.GitHubHost("https://github.com"),
			ClientID: os.Getenv("GITHUB_CLIENT_ID"),
			Scopes:   []string{"repo", "workflow"},
		}

		accessToken, err := flow.DetectFlow()
		if err != nil {
			return err
		}

		err = storage.StoreToken(storage.TokenKindGitHubGitea, accessToken)
		if err != nil {
			return err
		}

		fmt.Println("Successfully authenticated with GitHub")
	}

	return nil
}

// CreateRepo creates a repository on GitHub
func CreateRepo(name string) error {
	ctx := context.Background()
	storage.InitDefault()

	return createRepoWithContext(ctx, name)
}

// CreateRepoWithContext creates a repository on GitHub with a context
func createRepoWithContext(ctx context.Context, name string) error {
	accessToken := storage.GetToken(storage.TokenKindCLI)
	if accessToken == nil {
		fmt.Println("You are not authenticated, please run `gmg auth login` first.")
		return errors.New("not authenticated")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	org := os.Getenv("GITHUB_ORG")
	repo, _, err := client.Repositories.Create(ctx, org, &github.Repository{
		Name: &name,
	})
	if err != nil {
		fmt.Println("Failed to create repository:", err)
		return err
	}

	fmt.Println("Successfully created repository:", repo.GetCloneURL())

	return nil
}
