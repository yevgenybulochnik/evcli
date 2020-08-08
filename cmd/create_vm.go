package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var createVmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Create virtual machine",
	Long:  "Create virtual machine based on profile",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Please provide a vm name")
		}
		profile_flag, _ := cmd.Flags().GetString("profile")

		if !core.ProfileExists(profile_flag) {
			return errors.New("This profile does not exist")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// core.CreateVm()
		fmt.Println("test")
	},
}

func init() {
	createCmd.AddCommand(createVmCmd)
	createVmCmd.Flags().StringP("profile", "p", "", "Use configured profile")
	createVmCmd.MarkFlagRequired("profile")
}
