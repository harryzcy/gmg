package cmd

import (
	"fmt"
	"os"

	"github.com/harryzcy/gmg/internal/api"
	"github.com/harryzcy/gmg/internal/argutil"
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
		uri, err := argutil.GetURI(args)
		if err != nil {
			if err == argutil.ErrInvalidArgument {
				os.Exit(1)
			}
			fmt.Println(err)
			os.Exit(1)
		}

		options := api.MigrateRepoOptions{
			GitURI: uri,
		}

		options.Mirror, _ = cmd.Flags().GetBool("mirror")
		options.Name, _ = cmd.Flags().GetString("name")

		err = api.MigrateRepo(options)
		if err != nil {
			os.Exit(1)
		}
	},
}

var giteaMirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Mirror a repository to Gitea",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify the repository URL")
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Println("Please specify only one repository URL")
			os.Exit(1)
		}

		uri := args[0]
		options := api.MigrateRepoOptions{
			GitURI: uri,
			Mirror: true,
		}

		options.Name, _ = cmd.Flags().GetString("name")

		err := api.MigrateRepo(options)
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(giteaCmd)

	giteaCmd.AddCommand(giteaMigrateCmd)
	giteaMigrateCmd.Flags().BoolP("mirror", "m", false, "mirror the repository")
	giteaMigrateCmd.Flags().BoolP("private", "", false, "create the private repository")
	giteaMigrateCmd.Flags().StringP("name", "n", "", "name of the repository")

	giteaCmd.AddCommand(giteaMirrorCmd)
	giteaMirrorCmd.Flags().BoolP("private", "", false, "create the private repository")
	giteaMirrorCmd.Flags().StringP("name", "n", "", "name of the repository")
}
