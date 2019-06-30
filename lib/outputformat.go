package oremkrcli

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func OutputFormat(data [][]string, resourceType string) {
	table := tablewriter.NewWriter(os.Stdout)

	switch resourceType{
	case HOST:
		table.SetHeader([]string{"STATUS", "HOSTNAME", "ID",  "TYPE", "SERVICE/ROLE","CREATED"})
	}

	for _, value := range data{
		table.Append(value)
	}
	table.Render()
}