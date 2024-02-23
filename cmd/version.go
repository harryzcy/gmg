package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = ""
	commit  = ""
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(_ *cobra.Command, _ []string) {
		if version == "" {
			version = "dev"
		}
		if commit == "" {
			commit = "HEAD"
		}
		fmt.Printf("gmg version %s (commit %s)\n", version, commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
