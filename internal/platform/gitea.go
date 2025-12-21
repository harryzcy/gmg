package platform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"go.zcy.dev/gmg/internal/storage"
)

var (
	GiteaServer = os.Getenv("GITEA_SERVER")
	GiteaOrg    = os.Getenv("GITEA_ORG")
	GiteaToken  = os.Getenv("GITEA_ACCESS_TOKEN")
)

// MigrateRepoOptions represents the options for migrating a repository
type MigrateRepoOptions struct {
	GitURI  string
	Name    string
	Mirror  bool // if true, mirror the repository
	Private bool // if true, create a private repository
}

func MigrateRepo(options MigrateRepoOptions) error {
	if options.GitURI == "" {
		fmt.Println("HTTP(S) or git clone url is required")
		return errors.New("GitURI is required")
	}

	if options.Name == "" {
		options.Name = getNameFromGitURI(options.GitURI)
	}

	return requestMigration(options)
}

func getNameFromGitURI(gitURI string) string {
	parts := strings.Split(gitURI, "/")
	name := parts[len(parts)-1]

	name = strings.TrimSuffix(name, ".git")
	return name
}

// requestMigration requests a migration to Gitea
func requestMigration(options MigrateRepoOptions) (err error) {
	storage.InitDefault()

	token := storage.GetToken(storage.TokenKindGitHubGitea)
	if token == nil {
		fmt.Println("You are not authenticated, please run `gmg auth login` first.")
		return errors.New("GitHub token is required")
	}

	if GiteaServer == "" {
		fmt.Println("env GITEA_SERVER is required")
		return errors.New("env GITEA_SERVER is required")
	}

	url := GiteaServer + "/api/v1/repos/migrate"
	values := map[string]interface{}{
		"auth_token": token.Token,
		"clone_addr": options.GitURI,
		"mirror":     options.Mirror,
		"private":    options.Private,
		"repo_name":  options.Name,
		"repo_owner": GiteaOrg,
		"server":     "github",
	}
	data, err := json.Marshal(values)
	if err != nil {
		fmt.Println("Failed to marshal request body:", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+GiteaToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return err
	}
	defer func() {
		closeErr := resp.Body.Close()
		if err == nil {
			err = closeErr
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Failed to unmarshal response:", err)
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		err := fmt.Errorf("failed to create repository: %s", resp.Status)

		if resp.StatusCode == http.StatusConflict {
			fmt.Println("The repository with the same name already exists.")
			return err
		}

		if val, ok := result["message"]; ok {
			if message, ok := val.(string); ok {
				if strings.Contains(message, "user does not exist") {
					fmt.Println("Specified user does not exist, please check the GITEA_ORG environment variable.")
					return err
				}
			}
		}

		fmt.Println("Failed to create repository:", err)
		return err
	}

	fmt.Println("Successfully migrated repository:", options.Name)
	return nil
}

type SetupPushMirrorOptions struct {
	UsernameRepo string
	GitURI       string
}

func SetupPushMirror(options SetupPushMirrorOptions) error {
	storage.InitDefault()

	token := storage.GetToken(storage.TokenKindGitHubGitea)
	if token == nil {
		fmt.Println("You are not authenticated, please run `gmg auth login` first.")
		return errors.New("GitHub token is required")
	}

	if GiteaServer == "" {
		fmt.Println("env GITEA_SERVER is required")
		return errors.New("env GITEA_SERVER is required")
	}

	if ok := validateGitURI(options.GitURI); !ok {
		fmt.Println("invalid git uri")
		return errors.New("invalid git uri")
	}

	url := GiteaServer + "/api/v1/repos/" + options.UsernameRepo + "/push_mirrors"
	values := map[string]interface{}{
		"interval":        "0",
		"remote_address":  options.GitURI,
		"remote_password": token.Token,
		"remote_username": strings.SplitN(options.GitURI, "/", 3)[1],
		"sync_on_commit":  true,
	}
	data, err := json.Marshal(values)
	if err != nil {
		fmt.Println("Failed to marshal request body:", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+GiteaToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		err = fmt.Errorf("failed to create repository: %s", resp.Status)
		fmt.Println("Failed to create repository:", resp.Status)
		return err
	}

	fmt.Printf("Successfully added a push mirror to %s\n", options.UsernameRepo)
	return nil
}

func validateGitURI(gitURI string) bool {
	switch {
	case strings.HasPrefix(gitURI, "https://"):
		gitURI = strings.TrimPrefix(gitURI, "https://")
	case strings.HasPrefix(gitURI, "http://"):
		gitURI = strings.TrimPrefix(gitURI, "http://")
	case strings.HasPrefix(gitURI, "git@"):
		gitURI = strings.TrimPrefix(gitURI, "git@")
	default:
		return false
	}

	parts := strings.Split(gitURI, "/")
	if len(parts) != 3 {
		return false
	}

	if !strings.HasSuffix(parts[2], ".git") {
		return false
	}
	return true
}
