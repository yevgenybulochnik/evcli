package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var deletePoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "Delete storage pool",
	Long:  "Delete storage pool",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deletePool(args[0])
	},
}

func deletePool(poolName string) {
	conn := core.Connect()

	pools, err := conn.ListAllStoragePools(0)
	if err != nil {
		panic(err)
	}

	for _, pool := range pools {
		name, _ := core.GetPoolInfo(&pool)
		if name == poolName {
			accept := core.AskForConfirmation("Please confirm delete", 1)
			if accept {
				pool.Delete(0)
				pool.Undefine()
			} else {
				os.Exit(0)
			}
		}
	}

}

func init() {
	deleteCmd.AddCommand(deletePoolCmd)
}
