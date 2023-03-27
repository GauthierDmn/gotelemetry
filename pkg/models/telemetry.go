package models

import (
	"time"
)

type TelemetryType string

const (
	RealTime TelemetryType = "realtime"
	Playback TelemetryType = "playback"
)

type Telemetry struct {
	ReceivedTimestamp *time.Time
	Schema   *TelemetrySchema
	TelemetryType *TelemetryType
	Data     []byte
	Measurements []Measurement
}

func (t *Telemetry) GetMeasurement(name string) *Measurement {
	for _, measurement := range t.Measurements {
		if measurement.Field.Name == name {
			return &measurement
		}
	}

	return nil
}
