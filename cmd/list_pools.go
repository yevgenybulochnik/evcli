package cmd

import (

    "github.com/yevgenybulochnik/evcli/core"
    "github.com/spf13/cobra"
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
