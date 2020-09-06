package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var createNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Create network",
	Long:  "Create new libvirt network",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Please provide a newtwork name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cidrFlag, _ := cmd.Flags().GetString("cidr")
		core.CreateNetwork(args[0], cidrFlag)
	},
}

func init() {
	createCmd.AddCommand(createNetworkCmd)
	createNetworkCmd.Flags().StringP("cidr", "c", "", "Specifiy CIDR block")
	createNetworkCmd.MarkFlagRequired("cidr")
}
