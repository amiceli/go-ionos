package Ionos

import (
	"errors"
	"strings"
)

const apiBaseUrl = "https://api.hosting.ionos.com/"

type Api struct {
	ApiKey    string
	ApiSecret string
	Dns       Dns
}

func GetApi(apiKey string, apiSecret string) (*Api, error) {
	if len(apiKey) == 0 || len(apiSecret) == 0 {
		return nil, errors.New("missing env variables")
	}
	api := &Api{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}

	api.Dns = Dns{api: api}

	return api, nil
}

func (api *Api) apiKey() string {
	return strings.Join([]string{api.ApiKey, api.ApiSecret}, ".")
}

//

// type ZoneRecord struct {
// 	Id       string `json:"id"`
// 	Name     string `json:"name"`
// 	Type     string `json:"type"`
// 	RootName string `json:"rootName"`
// }

// type ZoneRecords struct {
// 	Recors []ZoneRecord `json:"records"`
// }

// type Zone struct {
// 	Id   string `json:"id"`
// 	Name string `json:"name"`
// 	Type string `json:"type"`
// }

// type ZoneList struct {
// 	Count int
// 	Zones []Zone
// }

// sturct avec attributs
// const (
// 	ZONES_API string = "dns/v1/zones"
// )

// func (zoneList *ZoneList) ZoneNames() []string {
// 	var names []string
// 	for _, domain := range zoneList.Zones {
// 		names = append(names, domain.Name)
// 	}

// 	return names
// }

// func (api *Api) LoadZones() ZoneList {
// 	url := fmt.Sprintf("%s/%s", apiBaseUrl, ZONES_API)

// 	client := http.Client{}
// 	req, _ := http.NewRequest("GET", url, nil)

// 	req.Header = http.Header{
// 		"accept":    {"application/json"},
// 		"X-Api-Key": {api.apiKey()},
// 	}

// 	res, _ := client.Do(req)

// 	zones, _ := Utils.ParseJSON[[]Zone](res.Body)

// 	return ZoneList{
// 		Count: len(zones),
// 		Zones: zones,
// 	}
// }

// func (api *Api) GetZone(zone Zone) ZoneRecords {
// 	url := fmt.Sprintf("%s/%s/%s?recordType=A", apiBaseUrl, ZONES_API, zone.Id)

// 	client := http.Client{}
// 	req, _ := http.NewRequest("GET", url, nil)

// 	req.Header = http.Header{
// 		"accept":    {"application/json"},
// 		"X-Api-Key": {api.apiKey()},
// 	}

// 	res, _ := client.Do(req)

// 	zoneDetails, _ := Utils.ParseJSON[ZoneRecords](res.Body)

// 	return zoneDetails
// }
