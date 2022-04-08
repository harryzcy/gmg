package storage

import (
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/cli/oauth/api"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	Init("./testdata")
	defer func() {
		err := os.RemoveAll("./testdata")
		if err != nil {
			t.Fatal(err)
		}
	}()

	tests := []struct {
		kind        TokenName
		token       *api.AccessToken
		expectedErr error
	}{
		{TokenKindCLI, &api.AccessToken{
			Token:        "cli-0token",
			RefreshToken: "cli-refresh-token",
			Type:         "cli-type",
			Scope:        "cli-scope",
		}, nil},
		{TokenKindGitHubGitea, &api.AccessToken{
			Token:        "g-token",
			RefreshToken: "g-refresh-token",
			Type:         "g-type",
			Scope:        "g-scope",
		}, nil},
		{TokenKindGitHubGitea, &api.AccessToken{
			Token:        "g-token-new",
			RefreshToken: "g-refresh-token-new",
			Type:         "g-type-new",
			Scope:        "g-scope-new",
		}, nil},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := StoreToken(test.kind, test.token)
			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.token, GetToken(test.kind))
		})
	}
}

func TestGetToken_Nil(t *testing.T) {
	Init("./testdata")
	defer func() {
		err := os.RemoveAll("./testdata")
		if err != nil {
			t.Fatal(err)
		}
	}()

	assert.Nil(t, GetToken("gh_test"))
}

func TestInitDefault(t *testing.T) {
	testDir := path.Join(os.Getenv("HOME"), ".gmg")

	baseDir = testDir
	InitDefault()
	assert.Equal(t, testDir, baseDir)
}
