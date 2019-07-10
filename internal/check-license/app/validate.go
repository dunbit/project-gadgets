package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate license headers",
	Run: runValidate,
}

func runValidate(cmd *cobra.Command, args []string) {
	fmt.Println("validate...")
}