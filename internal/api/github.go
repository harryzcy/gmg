package api

import (
	"fmt"
	"os"

	"github.com/cli/oauth"
)

func Auth() {
	flow := &oauth.Flow{
		Host:     oauth.GitHubHost("https://github.com"),
		ClientID: os.Getenv("OAUTH_CLIENT_ID"),
		Scopes:   []string{"repo"},
	}

	accessToken, err := flow.DetectFlow()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Access token: %s\n", accessToken.Token)
}
