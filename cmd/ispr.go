package cmd

import (
	"fmt"

	"github.com/klotzandrew/ci-info/checker"
	"github.com/spf13/cobra"
)

var isprCmd = &cobra.Command{
	Use:   "ispr",
	Short: "Detect if the current environment is a PR in CI",
	Long:  `Detect if the current environment is a PR in CI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(checker.IsPR())
	},
}

func init() {
	rootCmd.AddCommand(isprCmd)
}
