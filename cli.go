package main

import (
	"fmt"
	IonosApi "goolify/goolify-api"
	"os"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/jedib0t/go-pretty/v6/table"
)

func chooseZone(zoneList IonosApi.IonosZoneList) IonosApi.IonosZone {
	var names []string
	var selectedZone IonosApi.IonosZone

	for _, domain := range zoneList.Zones {
		names = append(names, domain.Name)
	}

	sp := selection.New("Which zone to load ?", names)
	choice, _ := sp.RunPrompt()

	for _, domain := range zoneList.Zones {
		if domain.Name == choice {
			selectedZone = domain
			break
		}
	}

	return selectedZone
}

func printRecords(records IonosApi.IonosZoneRecords) {
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
		api := IonosApi.Api{
			ApiKey:    envApi,
			ApiSecret: envSecret,
		}
		zoneList := api.LoadZones()

		fmt.Printf("%d zone(s) found", zoneList.Count)

		if zoneList.Count > 0 {
			fmt.Println("")
			fmt.Println("")

			selectedZone := chooseZone(zoneList)

			zoneRecords := api.GetZone(selectedZone)
			printRecords(zoneRecords)
		}
	}
}
