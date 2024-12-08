package Ionos

import (
	"fmt"
	Utils "go-ionos/utils"
	"net/http"
)

var RecordTypes = [20]string{
	"A", "AAAA", "CNAME", "MX", "NS", "SOA", "SRV", "TXT", "CAA", "TLSA", "SMIMEA", "SSHFP", "DS", "HTTPS", "SVCB", "CERT", "URI", "RP", "LOC", "OPENPGPKEY",
}

type Zone struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	// parent
	dns *Dns
}

type ZoneRecord struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	RootName string `json:"rootName"`
	Content  string `json:"content"`
	Ttl      int    `json:"ttl"`
	Prio     int    `json:"prio"`
	Disabled bool   `json:"disabled"`
}

type zoneRecords struct {
	Id     string       `json:"id"`
	Name   string       `json:"name"`
	Type   string       `json:"type"`
	Recors []ZoneRecord `json:"records"`
}

func (zone *Zone) GetRecords(recordType string) []ZoneRecord {
	var url string

	if recordType == "all" {
		url = fmt.Sprintf("%s/%s/%s", apiBaseUrl, ZONES_API, zone.Id)
	} else {
		url = fmt.Sprintf("%s/%s/%s?recordType=%s", apiBaseUrl, ZONES_API, zone.Id, recordType)
	}

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header = zone.dns.api.HttpHeaders()

	res, _ := client.Do(req)

	zoneDetails, _ := Utils.ParseJSON[zoneRecords](res.Body)

	return zoneDetails.Recors
}
