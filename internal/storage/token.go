package storage

import (
	"errors"
	"log"
	"os"
	"path"

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
		log.Fatal(err)
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
	TokenPrefix = "token."

	TokenKindCLI         TokenName = "gh_cli"
	TokenKindGitHubGitea TokenName = "gh_gitea"
)

func StoreToken(kind TokenName, token string) error {
	tokenViper.Set(TokenPrefix+string(kind), token)
	return saveTokenFile()
}

func GetToken(kind TokenName) string {
	return tokenViper.GetString(TokenPrefix + string(kind))
}

func saveTokenFile() error {
	file := path.Join(baseDir, "token.yaml")
	err := tokenViper.WriteConfigAs(file)
	return err
}
