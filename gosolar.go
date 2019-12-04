package gosolaredge

import (
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Handler struct {
	apiKey string
	siteID string
	apiURL string
	mqtt   mqtt.Client
}

var (
	apiURL = "https://monitoringapi.solaredge.com"
)

func New() *Handler {
	solarAPIKey, ok := os.LookupEnv("SOLAR_API_KEY")
	if !ok {
		panic("missing environment key: SOLAR_API_KEY")
	}

	solarSiteID, ok := os.LookupEnv("SOLAR_SITE_ID")
	if !ok {
		panic("missing environment key: SOLAR_SITE_ID")
	}

	log.Printf("SOLAR_API_KEY: %s*****", solarAPIKey[0:5])

	return &Handler{
		siteID: solarSiteID,
		apiKey: solarAPIKey,
		apiURL: apiURL,
	}
}
