package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: "create",
    Short: "Create libvirt resource",
    Long: "Create various libvirt resources including networks, pools and virtual machnines",
}

func init() {
	rootCmd.AddCommand(createCmd)
}
