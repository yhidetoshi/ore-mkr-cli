package oremkrcli

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

// OutputFormat output table.
func OutputFormat(data [][]string, resourceType string) error {
	table := tablewriter.NewWriter(os.Stdout)

	switch resourceType {
	case HOST:
		table.SetHeader([]string{"STATUS", "HOSTNAME", "ID", "TYPE", "SERVICE/ROLE", "CREATED"})
	case MONITOR:
		table.SetHeader([]string{"ID", "NAME", "SCOPE", "EXSCOPE", "WARNNING", "CRITICAL", "DURATION", "MAX_ATTEMPTS", "OVERVIEW"})
	case ALERT:
		table.SetHeader([]string{"ID", "NAME"})
	case USER:
		table.SetHeader([]string{"ID", "NAME", "AUTHORITY", "AUTHMETHOD", "MFA", "REGISTRATION", "JOINAT"})
	}

	for _, value := range data {
		table.Append(value)
	}
	table.Render()

	return nil
}
