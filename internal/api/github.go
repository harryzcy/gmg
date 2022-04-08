package api

import (
	"fmt"
	"os"

	"github.com/cli/oauth"
	"github.com/harryzcy/gmg/internal/storage"
)

func Auth() error {
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
			Scopes:   []string{"repo"},
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
