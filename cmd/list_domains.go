package cmd

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
	"libvirt.org/libvirt-go-xml"
)

var listVmsCmd = &cobra.Command{
	Use:   "vms",
	Short: "List vms",
	Long:  "List all virtual machines",
	Run: func(cmd *cobra.Command, args []string) {
		listVms()
	},
}

func listVms() {
	conn := core.Connect()

	domainsTable := CreateTableWriter(
		Header{"Name", "Status", "IPS", "Source"},
	)
	domainsTable.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Status", Align: text.AlignCenter},
		{Name: "IPS", Align: text.AlignCenter, AlignHeader: text.AlignCenter},
	})

	domains, _ := conn.ListAllDomains(0)
	for _, domain := range domains {
		domainXml, _ := domain.GetXMLDesc(0)
		d := &libvirtxml.Domain{}
		d.Unmarshal(domainXml)

		var currentState string

		switch status, _, _ := domain.GetState(); status {
		case 1:
			currentState = "up"
		case 5:
			currentState = "down"
		default:
			currentState = ""
		}

		interfaces, _ := domain.ListAllInterfaceAddresses(0)

		var ip string

		if len(interfaces) > 0 {
			ip = interfaces[0].Addrs[0].Addr
		} else {
			ip = ""
		}

		domainsTable.AppendRow(
			Row{d.Name, currentState, ip},
		)
	}

	domainsTable.Render()
}

func init() {
	listCmd.AddCommand(listVmsCmd)
}
