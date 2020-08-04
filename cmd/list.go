package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List various libvirt resources",
	Long:  `List libvirt resources including storage pools, networks and virtual machines`,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
