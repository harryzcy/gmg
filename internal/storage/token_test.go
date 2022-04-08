package storage

import (
	"os"
	"path"
	"strconv"
	"testing"

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
		token       string
		expectedErr error
	}{
		{TokenKindCLI, "gh_cli", nil},
		{TokenKindGitHubGitea, "gh_gitea", nil},
		{TokenKindGitHubGitea, "gh_gitea_new", nil},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := StoreToken(test.kind, test.token)
			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.token, GetToken(test.kind))
		})
	}
}

func TestInitDefault(t *testing.T) {
	testDir := path.Join(os.Getenv("HOME"), ".gmg")

	baseDir = testDir
	InitDefault()
	assert.Equal(t, testDir, baseDir)
}
