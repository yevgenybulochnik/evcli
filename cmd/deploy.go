package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy configuration",
	Long:  "Create multiple libvirt resources based on deployment file",
	Run: func(cmd *cobra.Command, args []string) {
		deploy()
	},
}

func deploy() {
	path, _ := os.Getwd()
	deploymentFile := filepath.Join(path, "deployment.yaml")

	data, _ := ioutil.ReadFile(deploymentFile)
	fmt.Println(deploymentFile)
	fmt.Println(data)

	var deployment core.Deployment

	deployment.Parse(data)

	globalProfiles, _ := core.GetGlobalProfiles()

	for name, host := range deployment.Hosts {
		if profile, exists := globalProfiles.List[host.Base]; exists {
			profile.CreateVM(name, 10, "vms")
		}
	}
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
