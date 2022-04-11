package cmd

import (
	"fmt"
	"os"

	"github.com/harryzcy/gmg/internal/api"
	"github.com/spf13/cobra"
)

// ghCmd represents the gh command
var ghCmd = &cobra.Command{
	Use:     "gh",
	Aliases: []string{"github"},
	Short:   "Manage GitHub repositories",
}

var repoCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a GitHub repository",
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
	rootCmd.AddCommand(ghCmd)

	ghCmd.AddCommand(repoCreateCmd)
}
