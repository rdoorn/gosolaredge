package gosolaredge

import (
	"encoding/json"
	"log"
	"testing"
)


func TestEquipment(t *testing.T) {
	data := []byte(`{"data":{"count":1,"telemetries":[{"date":"2019-12-04 13:14:12","totalActivePower":1534.0,"dcVoltage":747.812,"groundFaultResistance":7629.2,"powerLimit":100.0,"totalEnergy":25562.0,"temperature":39.4688,"inverterMode":"MPPT","operationMode":0,"vL1To2":390.906,"vL2To3":390.281,"vL3To1":393.875,"L1Data":{"acCurrent":2.26562,"acVoltage":227.078,"acFrequency":49.9939,"apparentPower":516.0,"activePower":512.0,"reactivePower":-64.0,"cosPhi":1.0},"L2Data":{"acCurrent":2.28125,"acVoltage":225.406,"acFrequency":49.9945,"apparentPower":516.0,"activePower":511.0,"reactivePower":-55.0,"cosPhi":1.0},"L3Data":{"acCurrent":2.27344,"acVoltage":225.984,"acFrequency":49.9945,"apparentPower":514.0,"activePower":511.0,"reactivePower":-71.0,"cosPhi":1.0}}]}}`)

	f := EquipmentDataHead{}

	json.Unmarshal(data, &f)
	log.Printf("solar: GET body response from api: %+v", f)
}
