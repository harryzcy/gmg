package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.zcy.dev/gmg/internal/platform"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with GitHub",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Obtain GitHub OAuth tokens for CLI and Gitea",
	Run: func(_ *cobra.Command, _ []string) {
		err := platform.Login()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.AddCommand(loginCmd)
}
