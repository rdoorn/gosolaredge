package gosolaredge

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
{"data":{"count":1,"telemetries":[{"date":"2019-12-01 11:04:10","totalActivePower":695.0,"dcVoltage":748.125,"groundFaultResistance":7629.17,"powerLimit":100.0,"totalEnergy":10112.0,"temperature":35.0047,"inverterMode":"MPPT","operationMode":0,
"vL1To2":395.531,"vL2To3":394.562,"vL3To1":398.688,"L1Data":{"acCurrent":1.13281,"acVoltage":230.219,"acFrequency":50.0101,"apparentPower":261.0,"activePower":234.0,"reactivePower":-115.0,"cosPhi":1.0},
"L2Data":{"acCurrent":1.10938,"acVoltage":228.188,"acFrequency":50.0082,"apparentPower":260.0,"activePower":228.0,"reactivePower":-107.0,"cosPhi":1.0},
"L3Data":{"acCurrent":1.13281,"acVoltage":228.141,"acFrequency":50.0098,"apparentPower":252.0,"activePower":233.0,"reactivePower":-115.0,"cosPhi":1.0}}]}}
*/

type EquipmentDataHead struct {
	EquipmentData EquipmentData `json:"data"`
}

type EquipmentData struct {
	Count       int         `json:"count"`
	Telemetries []Telemetry `json:"telemetries"`
}

type Telemetry struct {
	Date                  string  `json:"date"`
	TotalActivePower      float64 `json:"totalActivePower"`
	DCVoltate             float64 `json:"dcVoltage"`
	GroundFaultResistance float64 `json:"groundFaultResistance"`
	PowerLimit            float64 `json:"powerLimit"`
	TotalEnergy           float64 `json:"totalEnergy"`
	Temperature           float64 `json:"temperature"`
	VL1To2                float64 `json:"vL1To2"`
	VL2To3                float64 `json:"vL2To3"`
	VL3To1                float64 `json:"vL3To1"`
	L1Data                LData   `json:"L1Data"`
	L2Data                LData   `json:"L2Data"`
	L3Data                LData   `json:"L3Data"`
}

type LData struct {
	AcCurrent     float64 `json:"acCurrent"`
	AcVoltage     float64 `json:"acVoltage"`
	AcFrequency   float64 `json:"acFrequency"`
	ApparentPower float64 `json:"apparentPower"`
	ActivePower   float64 `json:"activePower"`
	ReactivePower float64 `json:"reactivePower"`
	CosPhi        float64 `json:"cosPhi"`
}

const TimeFormat = "2006-01-02 15:04:05"

func (h *Handler) ReadInventory(serialNR string, startTime, endTime time.Time) (EquipmentData, error) {

	url := fmt.Sprintf("%s/equipment/%s/%s/data.json?api_key=%s&startTime=%s&endTime=%s",
		h.apiURL, h.siteID, serialNR, h.apiKey, startTime.Format(TimeFormat), endTime.Format(TimeFormat))
	url = strings.Replace(url, " ", "%20", -1)

	log.Printf("request: %+v", url)
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return EquipmentData{}, err
	}

	defer res.Body.Close()

	data := EquipmentDataHead{}

	//bod, _ := ioutil.ReadAll(res.Body)
	//log.Printf("body: %s", bod)

	json.NewDecoder(res.Body).Decode(&data)
	log.Printf("solar: GET body response from api: %v", data)

	return data.EquipmentData, nil
}
