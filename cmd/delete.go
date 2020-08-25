package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete libvirt resource",
	Long:  "Delete various libvirt resources including networks, pools and virtual machines",
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
