package cmd

import (
	"net"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
	"libvirt.org/libvirt-go-xml"
)

var listNetworksCmd = &cobra.Command{
	Use:   "networks",
	Short: "List networks",
	Long:  `List all configured networks`,
	Run: func(cmd *cobra.Command, args []string) {
		listNetworks()
	},
}

func listNetworks() {
	conn := core.Connect()

	networksTable := CreateTableWriter(
		Header{"Name", "Mode", "Dhcp", "CIDR"},
	)

	networks, err := conn.ListAllNetworks(0)
	if err != nil {
		panic(err)
	}

	for _, network := range networks {
		networkXml, _ := network.GetXMLDesc(0)
		n := &libvirtxml.Network{}
		n.Unmarshal(networkXml)
		var cidr string

		if n.IPs[0].Prefix != 0 {
			_, ipNet, _ := net.ParseCIDR(n.IPs[0].Address + "/" + strconv.FormatUint(uint64(n.IPs[0].Prefix), 10))
			cidr = ipNet.String()
		} else {
			intMask, _ := net.IPMask(net.ParseIP(n.IPs[0].Netmask).To4()).Size()
			_, ipNet, _ := net.ParseCIDR(n.IPs[0].Address + "/" + strconv.Itoa(intMask))
			cidr = ipNet.String()
		}

		if n.IPs[0].DHCP != nil {
			networksTable.AppendRow(Row{n.Name, n.Forward.Mode, "true", cidr})
		} else {
			networksTable.AppendRow(Row{n.Name, n.Forward.Mode, "false", cidr})
		}
	}

	networksTable.Render()
}

func init() {
	listCmd.AddCommand(listNetworksCmd)
}
