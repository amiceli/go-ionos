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

	zones := api.Dns.GetZones()

	fmt.Printf("%d zone(s) found", len(zones))

	selectedZone := Cli.ChooseZone(zones)

	if selectedZone != nil {
		records := selectedZone.GetRecords()
		Cli.PrintRecords(records)
	}
}
