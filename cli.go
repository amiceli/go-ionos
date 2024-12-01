package main

import (
	"fmt"
	Ionos "go-ionos/ionos"
	"os"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/jedib0t/go-pretty/v6/table"
)

func printRecords(records Ionos.ZoneRecords) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Type"})

	for i, record := range records.Recors {
		t.AppendRow(
			table.Row{i, record.Name, record.Type},
		)
	}

	t.AppendFooter(table.Row{"", "Total", len(records.Recors)})
	t.Render()
}

func main() {
	envApi := os.Getenv("X-API-KEY")
	envSecret := os.Getenv("X-API-SECRET")

	if len(envApi) == 0 || len(envSecret) == 0 {
		fmt.Println("missing env variables")
	} else {
		api := Ionos.Api{
			ApiKey:    envApi,
			ApiSecret: envSecret,
		}
		zoneList := api.LoadZones()

		fmt.Printf("%d zone(s) found", zoneList.Count)

		if zoneList.Count > 0 {
			fmt.Println("")
			fmt.Println("")

			selectedZone := Ionos.ChooseZone(zoneList)

			if selectedZone != nil {
				zoneRecords := api.GetZone(*selectedZone)

				sp := selection.New("What do you want to do ?", []string{
					"show records", "add record", "remove record",
				})
				action, _ := sp.RunPrompt()

				switch action {
				case "show records":
					printRecords(zoneRecords)
				case "add record":
					break
				}
			}
		}
	}
}
