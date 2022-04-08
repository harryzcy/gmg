package storage

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/cli/oauth/api"
	"github.com/spf13/viper"
)

var (
	tokenViper *viper.Viper = viper.New()
	baseDir                 = path.Join(os.Getenv("HOME"), ".gmg")
)

func Init(directory string) {
	if directory != "" {
		baseDir = directory
	}
	err := loadToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func InitDefault() {
	Init("")
}

func loadToken() error {
	file := path.Join(baseDir, "token.yaml")
	tokenViper.SetConfigFile(file)
	err := tokenViper.ReadInConfig()
	if err != nil && errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(baseDir, os.ModePerm)
	}
	return err
}

type Tokens struct {
	GitHubToken string `json:"github_token"`
	GiteaToken  string `json:"gitea_token"`
}

type TokenName string

const (
	TokenKindCLI         TokenName = "gh_cli"
	TokenKindGitHubGitea TokenName = "gh_gitea"
)

func StoreToken(kind TokenName, token *api.AccessToken) error {
	name := string(kind)
	tokenViper.Set(name+".token", token.Token)
	tokenViper.Set(name+".refresh-token", token.RefreshToken)
	tokenViper.Set(name+".type", token.Type)
	tokenViper.Set(name+".scope", token.Scope)
	return saveTokenFile()
}

func GetToken(kind TokenName) *api.AccessToken {
	name := string(kind)
	return &api.AccessToken{
		Token:        tokenViper.GetString(name + ".token"),
		RefreshToken: tokenViper.GetString(name + ".refresh-token"),
		Type:         tokenViper.GetString(name + ".type"),
		Scope:        tokenViper.GetString(name + ".scope"),
	}
}

func saveTokenFile() error {
	file := path.Join(baseDir, "token.yaml")
	err := tokenViper.WriteConfigAs(file)
	return err
}
