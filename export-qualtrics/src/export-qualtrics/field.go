package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Field struct {
	SId        string           `json:"s-id,omitempty"`
	SType      string           `json:"s-type,omitempty"`
	SItems     []Field          `json:"s-items,omitempty"`
	SDirection string           `json:"s-direction,omitempty"`
	SGroup     string           `json:"s-group,omitempty"`
	SLabel     string           `json:"s-label,omitempty"`
	Html       string           `json:"html,omitempty"`
	Theme      string           `json:"data-theme,omitempty"`
	Value      string           `json:"value,omitempty"`
	Validation *FieldValidation `json:"validate,omitempty"`
}

func (field *Field) UnmarshalInterface(v interface{}) {
	if reflect.TypeOf(v).Kind() != reflect.Map {
		panic(fmt.Errorf("Cannot unmarshal value into Field: %s", v))
	}

	field_map := v.(map[string]interface{})

	unmarshalJSONString(&field.SId, field_map["s-id"])
	unmarshalJSONString(&field.SType, field_map["s-type"])
	unmarshalJSONFields(&field.SItems, field_map["s-items"])
	unmarshalJSONString(&field.SDirection, field_map["s-direction"])
	unmarshalJSONString(&field.SGroup, field_map["s-group"])
	unmarshalJSONString(&field.SLabel, field_map["s-label"])
	unmarshalJSONString(&field.Html, field_map["html"])
	unmarshalJSONString(&field.Theme, field_map["data-theme"])
	unmarshalJSONFieldValidation(&field.Validation, field_map["validate"])
}

func (field *Field) String() string {
	bytes, err := json.Marshal(field)
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

func (field *Field) UnmarshalKeyValueInterface(key string, value interface{}) {
	field.Html = key
	defer func() {
		if err := recover(); err != nil {
			if verbose {
				log.Printf("Unsupported subfield: %s\n", err)
			}
		}
	}()

	unmarshalJSONString(&field.Value, value)
}

type ByFieldValue []Field

func (f ByFieldValue) Len() int           { return len(f) }
func (f ByFieldValue) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByFieldValue) Less(i, j int) bool { return f[i].Value < f[j].Value }
