package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var deleteNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Delete network",
	Long:  "Delete network",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteNetwork(args[0])
	},
}

func deleteNetwork(netName string) {
	conn := core.Connect()

	networks, err := conn.ListAllNetworks(0)
	if err != nil {
		panic(err)
	}

	for _, network := range networks {
		name, _ := network.GetName()
		if name == netName {
			accept := core.AskForConfirmation("Please confirm delete", 1)
			if accept {
				network.Destroy()
				network.Undefine()
			} else {
				os.Exit(0)
			}
		}
	}

}

func init() {
	deleteCmd.AddCommand(deleteNetworkCmd)
}
