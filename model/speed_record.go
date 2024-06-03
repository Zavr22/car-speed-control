package models

import (
	"encoding/json"
	"time"
)

type SpeedRecord struct {
	Timestamp time.Time `json:"timestamp"`
	VehicleID string    `json:"vehicle_id"`
	Speed     float64   `json:"speed"`
}

const DateFormat = "02.01.2006 15:04:05"

func (r *SpeedRecord) UnmarshalJSON(data []byte) error {
	var aux struct {
		Timestamp string  `json:"timestamp"`
		VehicleID string  `json:"vehicle_id"`
		Speed     float64 `json:"speed"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	timestamp, err := time.Parse(DateFormat, aux.Timestamp)
	if err != nil {
		return err
	}
	r.Timestamp = timestamp
	r.VehicleID = aux.VehicleID
	r.Speed = aux.Speed
	return nil
}
