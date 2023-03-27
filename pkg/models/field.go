package models

import (
	"encoding/binary"
	"math"
)

type FieldType string

const (
	Uint8 FieldType = "uint8"
	Uint16 FieldType = "uint16"
	Uint32 FieldType = "uint32"
	Uint64 FieldType = "uint64"
	Int8 FieldType = "int8"
	Int16 FieldType = "int16"
	Int32 FieldType = "int32"
	Int64 FieldType = "int64"
	Float32 FieldType = "float32"
	Float64 FieldType = "float64"
	String FieldType = "string"
	Raw FieldType = "byte"
)

type Field struct {
	Name string
	Description string
	Type FieldType
	Size int
}

func (f *Field) GetValue(data []byte) interface{} {
	switch f.Type {
	case "uint8":
		return int(data[0])
	case "uint16":
		return binary.BigEndian.Uint16(data)
	case "uint32":
		return binary.BigEndian.Uint32(data)
	case "uint64":
		return binary.BigEndian.Uint64(data)
	case "int8":
		return int8(data[0])
	case "int16":
		return int16(binary.BigEndian.Uint16(data))
	case "int32":
		return int32(binary.BigEndian.Uint32(data))
	case "int64":
		return int64(binary.BigEndian.Uint64(data))
	case "float32":
		return math.Float32frombits(binary.BigEndian.Uint32(data))
	case "float64":
		return math.Float64frombits(binary.BigEndian.Uint64(data))
	case "string":
		return string(data)
	case "byte":
		return data
	default:
		return nil
	}
}
