package models

import (
	"testing"
	"math"
	"reflect"
)

func TestGetValue(t *testing.T) {
	const floatEpsilon = 0.000001
	testCases := []struct {
		name        string
		field       Field
		data        []byte
		expected    interface{}
	}{
		{
			name: "uint8",
			field: Field{
				Type: Uint8,
			},
			data:     []byte{0x01},
			expected: uint8(1),
		},
		{
			name: "uint16",
			field: Field{
				Type: Uint16,
			},
			data:     []byte{0x01, 0x02},
			expected: uint16(258),
		},
		{
			name: "uint32",
			field: Field{
				Type: Uint32,
			},
			data:     []byte{0x01, 0x02, 0x03, 0x04},
			expected: uint32(16909060),
		},
		{
			name: "uint64",
			field: Field{
				Type: Uint64,
			},
			data:     []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			expected: uint64(72623859790382856),
		},
		{
			name: "int8",
			field: Field{
				Type: Int8,
			},
			data:     []byte{0x7F},
			expected: int8(127),
		},
		{
			name: "int16",
			field: Field{
				Type: Int16,
			},
			data:     []byte{0x7F, 0xFF},
			expected: int16(32767),
		},
		{
			name: "int32",
			field: Field{
				Type: Int32,
			},
			data:     []byte{0x7F, 0xFF, 0xFF, 0xFF},
			expected: int32(2147483647),
		},
		{
			name: "int64",
			field: Field{
				Type: Int64,
			},
			data:     []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
			expected: int64(9223372036854775807),
		},
		{
			name: "float32",
			field: Field{
				Type: Float32,
			},
			data:     []byte{0x40, 0x49, 0x0F, 0xDB},
			expected: float32(3.141592),
		},
		{
			name: "float64",
			field: Field{
				Type: Float64,
			},
			data:     []byte{0x40, 0x09, 0x21, 0xFB, 0x54, 0x44, 0x2D, 0x18},
			expected: float64(3.141592653589793),
		},
		{
			name: "string",
			field: Field{
				Type: String,
			},
			data:     []byte("hello"),
			expected: "hello",
		},
		{
			name: "byte",
			field: Field{
				Type: Raw,
			},
			data:     []byte{0x01, 0x02, 0x03},
			expected: []byte{0x01, 0x02, 0x03},
		},
		{
			name: "invalid",
			field: Field{
				Type: "invalid",
			},
			data:     []byte{0x01},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.field.GetValue(tc.data)

			if vFloat, isFloat := value.(float64); isFloat {
				if eFloat, ok := tc.expected.(float64); ok {
					if math.Abs(vFloat-eFloat) > floatEpsilon {
						t.Errorf("expected: %v, got: %v", tc.expected, value)
					}
				} else {
					t.Errorf("value is float64, but expected value is not: %v", tc.expected)
				}
			} else if vFloat, isFloat := value.(float32); isFloat {
				if eFloat, ok := tc.expected.(float32); ok {
					if math.Abs(float64(vFloat-eFloat)) > floatEpsilon {
						t.Errorf("expected: %v, got: %v", tc.expected, value)
					}
				} else {
					t.Errorf("value is float32, but expected value is not: %v", tc.expected)
				}
			} else {
				if !reflect.DeepEqual(value, tc.expected) {
					t.Errorf("expected: %v, got: %v", tc.expected, value)
				}
			}
		})
	}
}
