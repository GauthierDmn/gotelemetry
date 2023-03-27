package models

import (
	"time"
)

type Measurement struct {
	Timestamp *time.Time
	RawValue  []byte
	Field    *Field
}

func (m *Measurement) GetValue() interface{} {
	return m.Field.GetValue(m.RawValue)
}
