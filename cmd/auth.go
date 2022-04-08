/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/harryzcy/gmg/internal/api"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with GitHub",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Obtain GitHub OAuth tokens for CLI and Gitea",
	Run: func(cmd *cobra.Command, args []string) {
		api.Login()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.AddCommand(loginCmd)
}
