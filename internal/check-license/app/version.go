package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/dunbit/project-gadgets/pkg/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of check-license",
	Run: runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Println(version.AppVersion)
}