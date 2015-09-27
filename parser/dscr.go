package parser

import (
	"bytes"
	"fmt"
)

type FieldName string

type Descriptor struct {
	fields   map[FieldName]interface{}
	AsString func(dscr *Descriptor, useAnsiColors bool) string
}

func (dscr Descriptor) AddField(fieldName FieldName, field interface{}) {
	dscr.fields[fieldName] = field
}

func (dscr Descriptor) SetString(fieldName FieldName, fieldValue string) {
	dscr.fields[fieldName] = fieldValue
}

func (dscr Descriptor) SetField(fieldName FieldName, fieldValue interface{}) {
	dscr.fields[fieldName] = fieldValue
}

func (dscr Descriptor) AppendItem(fieldName FieldName, fieldValue interface{}) {
	dscr.fields[fieldName] = append(dscr.fields[fieldName].([]interface{}), fieldValue)
}

func (dscr Descriptor) SetMapItem(fieldName FieldName, key string, value interface{}) {
	dscr.fields[fieldName].(map[string]interface{})[key] = value
}

func (dscr Descriptor) GetField(fieldName FieldName) interface{} {
	return dscr.fields[fieldName]
}

func (dscr Descriptor) GetMapItem(fieldName FieldName, key string) interface{} {
	return dscr.fields[fieldName].(map[string]interface{})[key]
}

func (dscr Descriptor) GetString(fieldName FieldName) string {
	return dscr.fields[fieldName].(string)
}

func (dscr Descriptor) GetInt(fieldName FieldName) int {
	return dscr.fields[fieldName].(int)
}

func (dscr Descriptor) GetSlice(fieldName FieldName) []interface{} {
	return dscr.fields[fieldName].([]interface{})
}

func (dscr Descriptor) AddToOut(buffer *bytes.Buffer, output string, color string) {
	if color == "" {
		fmt.Fprintf(buffer, output)
	} else {
		lines := bytes.Split([]byte(output), []byte{'\n'})
		if len(lines) > 1 {
			for _, line := range lines {
				if string(line) != "" {
					fmt.Fprintf(buffer, "\033[%sm%s\033[0m\n", color, string(line))
				}
			}
		} else {
			fmt.Fprintf(buffer, "\033[%sm%s\033[0m", color, output)
		}
	}
}

func (dscr Descriptor) String() string {
	return dscr.AsString(&dscr, false)
}

func NewDescriptor() *Descriptor {
	dscr := new(Descriptor)
	dscr.fields = map[FieldName]interface{}{}
	dscr.AsString = func(d *Descriptor, useAnsiColors bool) string {
		output := new(bytes.Buffer)
		for name, value := range d.fields {
			fmt.Fprintf(output, "%s=%s\n", name, value)
		}
		return output.String()
	}
	return dscr
}
