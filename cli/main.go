package cli

import (
	"fmt"
	Ionos "go-ionos/ionos"
	"os"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/jedib0t/go-pretty/v6/table"
)

func ChooseZone(zones []Ionos.Zone) *Ionos.Zone {
	var names []string
	for _, zone := range zones {
		names = append(names, zone.Name)
	}

	sp := selection.New("Select a zone", names)
	choice, _ := sp.RunPrompt()

	for _, zone := range zones {
		if zone.Name == choice {
			return &zone
		}
	}

	return nil
}

func ChooseRecordTpe() string {
	allTypes := append(Ionos.RecordTypes[:], "all")

	sp := selection.New("Select record type", allTypes)
	sp.PageSize = 10
	choice, _ := sp.RunPrompt()

	return choice
}

func PrintRecords(records []Ionos.ZoneRecord, recordType string) {
	if len(records) == 0 {
		fmt.Printf("No %s record found", recordType)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Type"})

	for i, record := range records {
		t.AppendRow(
			table.Row{i, record.Name, record.Type},
		)
	}

	t.AppendFooter(table.Row{"", "Total", len(records)})
	t.Render()
}
