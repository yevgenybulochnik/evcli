package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var listNetworksCmd = &cobra.Command{
	Use: "networks",
	Run: func(cmd *cobra.Command, args []string) {
		core.ListNetworks()
	},
}

func init() {
	listCmd.AddCommand(listNetworksCmd)
}
