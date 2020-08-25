package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
	"libvirt.org/libvirt-go-xml"
)

var createPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "Create storage pool",
	Long:  "Create libvirt storage pool based on path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pathFlag, _ := cmd.Flags().GetString("path")
		createPool(args[0], pathFlag)
	},
}

func createPool(name string, path string) {
	conn := core.Connect()
	poolcfg := libvirtxml.StoragePool{
		Name: name,
		Type: "dir",
		Target: &libvirtxml.StoragePoolTarget{
			Path: path,
		},
	}

	xml, _ := poolcfg.Marshal()

	pool, _ := conn.StoragePoolDefineXML(xml, 0)

	pool.Create(0)
	pool.SetAutostart(true)
}

func init() {
	createCmd.AddCommand(createPoolCmd)
	createPoolCmd.Flags().StringP("path", "p", "", "Storage pool path")
	createPoolCmd.MarkFlagRequired("path")
}
