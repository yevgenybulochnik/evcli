package core

import (
	"encoding/xml"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	libvirt "libvirt.org/libvirt-go"
)

type Pool struct {
	XMLName xml.Name `xml:"pool"`
	Name    string   `xml:"name"`
	Path    string   `xml:"target>path"`
}

type header = table.Row

type row = table.Row

func createTableWriter(h header) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(h)
	return t
}

func ListPools() {
	conn := Connect()

	poolsTable := createTableWriter(
		header{"Pool", "Path"},
	)

	pools, err := conn.ListAllStoragePools(libvirt.CONNECT_LIST_STORAGE_POOLS_ACTIVE)
	if err != nil {
		fmt.Print(err)
	}

	for _, pool := range pools {
		poolXml, _ := pool.GetXMLDesc(0)
		var p Pool
		xml.Unmarshal([]byte(poolXml), &p)
		poolsTable.AppendRow(row{p.Name, p.Path})
	}

	poolsTable.Render()
}

type Network struct {
	XMLName xml.Name `xml:"network"`
	Name    string   `xml:"name"`
	Forward struct {
		Mode string `xml:"mode,attr"`
	} `xml:"forward"`
	Ip struct {
		Dhcp    string `xml:"dhcp"`
		Address string `xml:"address,attr"`
		Prefix  string `xml:"prefix,attr"`
		Mask    string `xml:"netmask,attr"`
	} `xml:"ip"`
}

func ListNetworks() {
	conn := Connect()

	networksTable := createTableWriter(
		header{"Name", "Mode", "Dhcp", "CIDR"},
	)

	networks, err := conn.ListAllNetworks(0)
	if err != nil {
		fmt.Print(err)
	}

	for _, network := range networks {
		networkXml, _ := network.GetXMLDesc(0)
		var n Network
		xml.Unmarshal([]byte(networkXml), &n)
		var cidr string

		if n.Ip.Prefix != "" {
			_, ipNet, _ := net.ParseCIDR(n.Ip.Address + "/" + n.Ip.Prefix)
			cidr = ipNet.String()
		} else {
			intMask, _ := net.IPMask(net.ParseIP(n.Ip.Mask).To4()).Size()
			_, ipNet, _ := net.ParseCIDR(n.Ip.Address + "/" + strconv.Itoa(intMask))
			cidr = ipNet.String()
		}

		if n.Ip.Dhcp != "" {
			networksTable.AppendRow(row{n.Name, n.Forward.Mode, "true", cidr})
		} else {
			networksTable.AppendRow(row{n.Name, n.Forward.Mode, "false", cidr})
		}
	}

	networksTable.Render()
}
