package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
	"gopkg.in/yaml.v3"
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
		profilesFile, _ := core.GetProfilesFile()

		var profileConfig core.ProfileConfig

		yaml.Unmarshal(profilesFile, &profileConfig)
		p, _ := cmd.Flags().GetString("profile")
		profile := profileConfig.Profiles[p]
		backingImage := profile.GetImagePath()
		core.CreateImage(args[0], backingImage, 20, "/home/yevgeny/vms")
		core.CreateVm(args[0], "/home/yevgeny/vms")

	},
}

func init() {
	createCmd.AddCommand(createVmCmd)
	createVmCmd.Flags().StringP("profile", "p", "", "Use configured profile")
	createVmCmd.MarkFlagRequired("profile")
}
