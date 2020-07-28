package cmd

import (
	"github.com/spf13/cobra"
)

var createVmCmd = &cobra.Command{
	Use: "vm",
}

func init() {
	createCmd.AddCommand(createVmCmd)
}
