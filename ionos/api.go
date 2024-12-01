package Ionos

import (
	"fmt"
	Utils "go-ionos/utils"
	"net/http"
	"strings"
)

type Api struct {
	ApiKey    string
	ApiSecret string
}

type ZoneRecord struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	RootName string `json:"rootName"`
}

type ZoneRecords struct {
	Recors []ZoneRecord `json:"records"`
}

type Zone struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ZoneList struct {
	Count int
	Zones []Zone
}

func (zoneList *ZoneList) ZoneNames() []string {
	var names []string
	for _, domain := range zoneList.Zones {
		names = append(names, domain.Name)
	}

	return names
}

func (api *Api) BaseUrl() string {
	return "https://api.hosting.ionos.com/"
}

func (api *Api) apiKey() string {
	return strings.Join([]string{api.ApiKey, api.ApiSecret}, ".")
}

func (api *Api) LoadZones() ZoneList {
	url := fmt.Sprintf("%s/dns/v1/zones", api.BaseUrl())

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header = http.Header{
		"accept":    {"application/json"},
		"X-Api-Key": {api.apiKey()},
	}

	res, _ := client.Do(req)

	zones, _ := Utils.ParseJSON[[]Zone](res.Body)

	return ZoneList{
		Count: len(zones),
		Zones: zones,
	}
}

func (api *Api) GetZone(zone Zone) ZoneRecords {
	url := fmt.Sprintf("%s/dns/v1/zones/%s?recordType=A", api.BaseUrl(), zone.Id)

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header = http.Header{
		"accept":    {"application/json"},
		"X-Api-Key": {api.apiKey()},
	}

	fmt.Sprintln("Before")
	res, _ := client.Do(req)

	zoneDetails, _ := Utils.ParseJSON[ZoneRecords](res.Body)

	return zoneDetails
}
