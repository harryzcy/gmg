package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.zcy.dev/gmg/internal/api"
	"go.zcy.dev/gmg/internal/argutil"
)

// giteaCmd represents the gitea command
var giteaCmd = &cobra.Command{
	Use:   "gitea",
	Short: "Manage Gitea repositories",
}

var giteaMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Args:  cobra.ExactArgs(1),
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
		options.Private, _ = cmd.Flags().GetBool("private")

		err = api.MigrateRepo(options)
		if err != nil {
			os.Exit(1)
		}
	},
}

var giteaMirrorCmd = &cobra.Command{
	Use:   "mirror url",
	Args:  cobra.ExactArgs(1),
	Short: "Mirror a repository to Gitea",
	Run: func(cmd *cobra.Command, args []string) {
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

var giteaPushMirrorCmd = &cobra.Command{
	Use:     "pushmirror",
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"pm"},
	Short:   "Setup a push mirror to a repository",
	Run: func(cmd *cobra.Command, args []string) {
		options := api.SetupPushMirrorOptions{
			UsernameRepo: args[0],
			GitURI:       args[1],
		}

		err := api.SetupPushMirror(options)
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

	giteaCmd.AddCommand(giteaPushMirrorCmd)
}
