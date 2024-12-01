package IonosApi

import (
	"fmt"
	"net/http"
	"strings"
)

type Api struct {
	ApiKey    string
	ApiSecret string
}

type IonosZoneRecord struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	RootName string `json:"rootName"`
}

type IonosZoneRecords struct {
	Recors []IonosZoneRecord `json:"records"`
}

type IonosZone struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type IonosZoneList struct {
	Count int
	Zones []IonosZone
}

func (api *Api) BaseUrl() string {
	return "https://api.hosting.ionos.com/"
}

func (api *Api) apiKey() string {
	return strings.Join([]string{api.ApiKey, api.ApiSecret}, ".")
}

func (api *Api) LoadZones() IonosZoneList {
	url := fmt.Sprintf("%s/dns/v1/zones", api.BaseUrl())

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header = http.Header{
		"accept":    {"application/json"},
		"X-Api-Key": {api.apiKey()},
	}

	res, _ := client.Do(req)

	zones, _ := ParseJSON[[]IonosZone](res.Body)

	return IonosZoneList{
		Count: len(zones),
		Zones: zones,
	}
}

func (api *Api) GetZone(zone IonosZone) IonosZoneRecords {
	url := fmt.Sprintf("%s/dns/v1/zones/%s?recordType=A", api.BaseUrl(), zone.Id)

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header = http.Header{
		"accept":    {"application/json"},
		"X-Api-Key": {api.apiKey()},
	}

	fmt.Sprintln("Before")
	res, _ := client.Do(req)

	zoneDetails, _ := ParseJSON[IonosZoneRecords](res.Body)

	return zoneDetails
}
