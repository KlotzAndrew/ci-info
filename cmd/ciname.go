package cmd

import (
	"fmt"

	"github.com/klotzandrew/ci-info/ci"
	"github.com/spf13/cobra"
)

var cinameCmd = &cobra.Command{
	Use:   "ciname",
	Short: "Return name of CI",
	Long:  `Return name of CI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ci.CIName())
	},
}

func init() {
	rootCmd.AddCommand(cinameCmd)
}
