package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var listVmsCmd = &cobra.Command{
	Use: "vms",
	Run: func(cmd *cobra.Command, args []string) {
		core.ListDomains()
	},
}

func init() {
	listCmd.AddCommand(listVmsCmd)
}
