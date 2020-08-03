package cmd

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
	"libvirt.org/libvirt-go-xml"
)

var listPoolsCmd = &cobra.Command{
	Use:   "pools",
	Short: "List storage pools",
	Long:  `List all configured storage pools`,
	Run: func(cmd *cobra.Command, args []string) {
		listPools()
	},
}

func listPools() {
	conn := core.Connect()

	poolsTable := CreateTableWriter(
		Header{"Pool", "Path"},
	)
	poolsTable.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Path", AlignHeader: text.AlignCenter},
	})

	pools, err := conn.ListAllStoragePools(0)
	if err != nil {
		panic(err)
	}

	for _, pool := range pools {
		poolXml, _ := pool.GetXMLDesc(0)
		p := &libvirtxml.StoragePool{}
		p.Unmarshal(poolXml)
		poolsTable.AppendRow(Row{p.Name, p.Target.Path})
	}

	poolsTable.Render()
}
