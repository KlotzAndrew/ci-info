package cmd

import (
	"fmt"

	"github.com/klotzandrew/ci-info/ci"
	"github.com/spf13/cobra"
)

// cancheckprCmd represents the cancheckpr command
var cancheckprCmd = &cobra.Command{
	Use:   "cancheckpr",
	Short: "Return if PR check is supported for this vendor",
	Long:  `Return if PR check is supported for this vendor`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ci.CanCheckPR())
	},
}

func init() {
	rootCmd.AddCommand(cancheckprCmd)
}
