package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.zcy.dev/gmg/internal/argutil"
	"go.zcy.dev/gmg/internal/platform"
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

		options := platform.MigrateRepoOptions{
			GitURI: uri,
		}

		options.Mirror, err = cmd.Flags().GetBool("mirror")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		options.Name, err = cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		options.Private, err = cmd.Flags().GetBool("private")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = platform.MigrateRepo(options)
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
		options := platform.MigrateRepoOptions{
			GitURI: uri,
			Mirror: true,
		}

		var err error
		options.Name, err = cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = platform.MigrateRepo(options)
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
	Run: func(_ *cobra.Command, args []string) {
		options := platform.SetupPushMirrorOptions{
			UsernameRepo: args[0],
			GitURI:       args[1],
		}

		err := platform.SetupPushMirror(options)
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
