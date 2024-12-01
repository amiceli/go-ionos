package Ionos

import "github.com/erikgeiser/promptkit/selection"

func ChooseZone(zoneList ZoneList) *Zone {
	sp := selection.New("Which zone to manage ?", zoneList.ZoneNames())
	choice, _ := sp.RunPrompt()

	for _, zone := range zoneList.Zones {
		if zone.Name == choice {
			return &zone
		}
	}

	return nil
}
