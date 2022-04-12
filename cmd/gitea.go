package cmd

import (
	"fmt"
	"os"

	"github.com/harryzcy/gmg/internal/api"
	"github.com/spf13/cobra"
)

// giteaCmd represents the gitea command
var giteaCmd = &cobra.Command{
	Use:   "gitea",
	Short: "Manage Gitea repositories",
}

var giteaMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate a repository to Gitea",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the repository URL")
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Println("Please specify only one repository URL")
			os.Exit(1)
		}

		mirror, err := cmd.Flags().GetBool("mirror")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		uri := args[0]
		err = api.MigrateRepo(api.MigrateRepoOptions{
			GitURI: uri,
			Mirror: mirror,
		})
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(giteaCmd)

	giteaCmd.AddCommand(giteaMigrateCmd)
	giteaMigrateCmd.Flags().BoolP("mirror", "m", false, "mirror the repository")
}
