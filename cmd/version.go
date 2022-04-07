package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = ""
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gmg version", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
