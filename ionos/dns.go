package Ionos

import (
	"fmt"
	Utils "go-ionos/utils"
	"net/http"
)

const (
	ZONES_API string = "dns/v1/zones"
)

type Dns struct {
	api *Api
}

func (dns *Dns) GetZones() []Zone {
	url := fmt.Sprintf("%s/%s", apiBaseUrl, ZONES_API)

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header = dns.api.HttpHeaders()

	res, _ := client.Do(req)

	zones, _ := Utils.ParseJSON[[]Zone](res.Body)

	for i := range zones {
		zones[i].dns = dns
	}

	return zones
}
