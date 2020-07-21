package core

import (
	"fmt"
	"net"
	"strconv"

	"libvirt.org/libvirt-go"
	"libvirt.org/libvirt-go-xml"
)

func ListPools() {
	conn := Connect()

	poolsTable := CreateTableWriter(
		Header{"Pool", "Path"},
	)

	pools, err := conn.ListAllStoragePools(libvirt.CONNECT_LIST_STORAGE_POOLS_ACTIVE)
	if err != nil {
		fmt.Print(err)
	}

	for _, pool := range pools {
		poolXml, _ := pool.GetXMLDesc(0)
		p := &libvirtxml.StoragePool{}
		p.Unmarshal(poolXml)
		poolsTable.AppendRow(Row{p.Name, p.Target.Path})
	}

	poolsTable.Render()
}

func ListNetworks() {
	conn := Connect()

	networksTable := CreateTableWriter(
		Header{"Name", "Mode", "Dhcp", "CIDR"},
	)

	networks, err := conn.ListAllNetworks(0)
	if err != nil {
		fmt.Print(err)
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
