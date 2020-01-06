package cmd

import (
	"fmt"

	"github.com/klotzandrew/ci-info/ci"
	"github.com/spf13/cobra"
)

var isciCmd = &cobra.Command{
	Use:   "isci",
	Short: "Detect if the current environment is CI",
	Long:  `Detect if the current environment is CI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ci.IsCI())
	},
}

func init() {
	rootCmd.AddCommand(isciCmd)
}
