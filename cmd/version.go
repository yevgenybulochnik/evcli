package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of evcli",
	Long:  "Version number for evcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version test")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
