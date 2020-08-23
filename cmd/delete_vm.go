package cmd

import (
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
	"libvirt.org/libvirt-go-xml"
)

var deleteVmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Create virtual machine",
	Long:  "Create virtual machine based on profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteVm(args[0])
	},
}

func init() {
	deleteCmd.AddCommand(deleteVmCmd)
}

func deleteVm(vmName string) {
	conn := core.Connect()

	dom, err := conn.LookupDomainByName(vmName)
	domXmlData, _ := dom.GetXMLDesc(0)
	domXml := &libvirtxml.Domain{}
	domXml.Unmarshal(domXmlData)

	pools, _ := conn.ListAllStoragePools(0)
	for _, pool := range pools {
		pool.Refresh(0)
	}

	volPath := domXml.Devices.Disks[0].Source.File.File
	isoPath := strings.TrimSuffix(volPath, filepath.Ext(volPath)) + ".iso"

	iso, _ := conn.LookupStorageVolByPath(isoPath)
	vol, _ := conn.LookupStorageVolByPath(volPath)

	dom.Destroy()
	dom.Undefine()
	iso.Delete(0)
	vol.Delete(0)

	if err != nil {
		panic(err)
	}

}
