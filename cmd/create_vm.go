package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var createVmCmd = &cobra.Command{
	Use: "vm",
	Run: func(cmd *cobra.Command, args []string) {
		core.CreateVm()
	},
}

func init() {
	createCmd.AddCommand(createVmCmd)
}
