package Ionos

import (
	"fmt"
	Utils "go-ionos/utils"
	"net/http"
)

type Zone struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	// parent
	Dns *Dns
}

type ZoneRecord struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	RootName string `json:"rootName"`
}

type zoneRecords struct {
	Recors []ZoneRecord `json:"records"`
}

func (zone *Zone) GetRecords() []ZoneRecord {
	url := fmt.Sprintf("%s/%s/%s?recordType=A", apiBaseUrl, ZONES_API, zone.Id)

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header = http.Header{
		"accept":    {"application/json"},
		"X-Api-Key": {zone.Dns.api.apiKey()},
	}

	res, _ := client.Do(req)

	zoneDetails, _ := Utils.ParseJSON[zoneRecords](res.Body)

	return zoneDetails.Recors
}
