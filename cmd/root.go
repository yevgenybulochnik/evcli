package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "evcli",
	Short: "Easy Virt Cli",
	Long: `Go wrapper around libvirt to help with easy creation of vms.
Storage pools and networks can also be configured!
  `,
}

func Execute() error {
	return rootCmd.Execute()
}
