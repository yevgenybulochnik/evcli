package cmd

import (

    "github.com/yevgenybulochnik/evcli/core"
    "github.com/spf13/cobra"
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
