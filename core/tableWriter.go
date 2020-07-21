package core

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Header = table.Row

type Row = table.Row

func CreateTableWriter(h Header) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(h)
	return t
}
