package cli

import (
	"fmt"
	Ionos "go-ionos/ionos"
	"os"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/jedib0t/go-pretty/v6/table"
)

const (
	LIST_ZONE         = "List zones"
	LIST_ZONE_RECORDS = "Get zone records"
)

var commands = []string{LIST_ZONE, LIST_ZONE_RECORDS}

func ChooseCommand() string {
	sp := selection.New("What do you want to do", commands)
	choice, _ := sp.RunPrompt()

	for _, command := range commands {
		if command == choice {
			return command
		}
	}

	return ""
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

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

type PrintRecordsOptions struct {
	Records    []Ionos.ZoneRecord
	RecordType string
	Zone       Ionos.Zone
}

func PrintRecords(options PrintRecordsOptions) {
	if len(options.Records) == 0 {
		fmt.Printf("No %s record found in %s", options.RecordType, options.Zone.Name)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Type", "Content"})

	for i, record := range options.Records {
		t.AppendRow(
			table.Row{i, record.Name, record.Type, record.Content},
		)
	}

	t.AppendFooter(table.Row{"", "Total", len(options.Records)})
	t.Render()
}

func PrintZones(zones []Ionos.Zone) {
	if len(zones) == 0 {
		fmt.Println("No zones found")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Id", "Name", "Type"})

	for i, zone := range zones {
		t.AppendRow(
			table.Row{i, zone.Id, zone.Name, zone.Type},
		)
	}

	t.AppendFooter(table.Row{"", "Total", "", len(zones)})
	t.Render()
}
