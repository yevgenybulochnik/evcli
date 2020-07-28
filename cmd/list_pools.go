package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var listPoolsCmd = &cobra.Command{
	Use: "pools",
	Run: func(cmd *cobra.Command, args []string) {
		core.ListPools()
	},
}

func init() {
	listCmd.AddCommand(listPoolsCmd)
}
