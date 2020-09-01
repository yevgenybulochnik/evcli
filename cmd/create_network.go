package cmd

import (
	"github.com/spf13/cobra"
)

var createNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Create network",
	Long:  "Create new libvirt network",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func createNetwork() {
}

func init() {
	createCmd.AddCommand(createNetworkCmd)
}
