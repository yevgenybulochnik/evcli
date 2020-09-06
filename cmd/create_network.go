package cmd

import (
	"errors"
	"fmt"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
	"libvirt.org/libvirt-go-xml"
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
		createNetwork(args[0], cidrFlag)
	},
}

func createNetwork(netName string, cidrString string) {
	conn := core.Connect()
	_, ipNet, _ := net.ParseCIDR(cidrString)
	count := cidr.AddressCount(ipNet)
	if count <= 4 {
		panic("Error calculating DHCP range")
	}
	firstIp, lastIp := cidr.AddressRange(ipNet)
	netcfg := &libvirtxml.Network{
		Name: netName,
		Forward: &libvirtxml.NetworkForward{
			Mode: "nat",
		},
		Bridge: &libvirtxml.NetworkBridge{
			Name:  netName,
			STP:   "on",
			Delay: "0",
		},
		IPs: []libvirtxml.NetworkIP{
			{
				Address: cidr.Inc(firstIp).String(),
				Netmask: ipv4MaskString(ipNet.Mask),
				DHCP: &libvirtxml.NetworkDHCP{
					Ranges: []libvirtxml.NetworkDHCPRange{
						{
							Start: cidr.Inc(cidr.Inc(firstIp)).String(),
							End:   cidr.Dec(lastIp).String(),
						},
					},
				},
			},
		},
	}

	xml, _ := netcfg.Marshal()

	net, err := conn.NetworkDefineXML(xml)
	if err != nil {
		panic(err)
	}
	net.Create()
	net.SetAutostart(true)

	fmt.Println(netcfg.Marshal())
	fmt.Println(cidr.AddressRange(ipNet))
}

func ipv4MaskString(m []byte) string {
	if len(m) != 4 {
		panic("ipv4 mask must have a length of 4 bytes")
	}

	return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
}

func init() {
	createCmd.AddCommand(createNetworkCmd)
	createNetworkCmd.Flags().StringP("cidr", "c", "", "Specifiy CIDR block")
	createNetworkCmd.MarkFlagRequired("cidr")
}
