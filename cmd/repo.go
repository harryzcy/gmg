/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/harryzcy/gmg/internal/api"
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage repositories",
}

var repoCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a repository name")
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Println("Please specify only one repository name")
			os.Exit(1)
		}

		name := args[0]
		err := api.CreateRepo(name)
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)

	repoCmd.AddCommand(repoCreateCmd)
}
