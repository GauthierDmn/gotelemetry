package models

type TelemetrySchema struct {
	Name string
	Description string
	Fields []Field
}

func (s *TelemetrySchema) GetField(name string) *Field {
	for _, field := range s.Fields {
		if field.Name == name {
			return &field
		}
	}

	return nil
}

func (s *TelemetrySchema) GetFields() []Field {
	return s.Fields
}

func (s *TelemetrySchema) GetFieldsCount() int {
	return len(s.Fields)
}

func (s *TelemetrySchema) GetFieldAt(index int) *Field {
	return &s.Fields[index]
}

func (s *TelemetrySchema) GetSize() int {
	size := 0

	for _, field := range s.Fields {
		size += field.Size
	}

	return size
}

func (s *TelemetrySchema) isCompatibleWithData(data []byte) bool {
	return len(data) == s.GetSize()
}
