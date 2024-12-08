package main

import (
	"fmt"
	Cli "go-ionos/cli"
	Ionos "go-ionos/ionos"
	"os"
)

func main() {
	envApi := os.Getenv("X-API-KEY")
	envSecret := os.Getenv("X-API-SECRET")

	api, err := Ionos.GetApi(envApi, envSecret)

	if err != nil {
		fmt.Println("missing env variables")
		return
	}

	Cli.ClearTerminal()
	var command = Cli.ChooseCommand()

	switch command {
	case Cli.LIST_ZONE:
		zones := api.Dns.GetZones()
		Cli.PrintZones(zones)
	case Cli.LIST_ZONE_RECORDS:
		zones := api.Dns.GetZones()
		selectedZone := Cli.ChooseZone(zones)

		if selectedZone != nil {
			recordType := Cli.ChooseRecordTpe()
			records := selectedZone.GetRecords(recordType)

			Cli.PrintRecords(Cli.PrintRecordsOptions{
				Zone:       *selectedZone,
				RecordType: recordType,
				Records:    records,
			})
		}
	default:
		return
	}
}
