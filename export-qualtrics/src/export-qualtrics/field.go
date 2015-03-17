package main

import (
	"encoding/json"
	"log"
	"strconv"
)

type Field struct {
	SId        string           `json:"s-id"`
	SType      string           `json:"s-type"`
	SItems     SItems           `json:"s-items"`
	SDirection string           `json:"direction"`
	SGroup     string           `json:"s-group"`
	SLabel     string           `json:"s-label"`
	Html       string           `json:"html"`
	Theme      string           `json:"data-theme"`
	Validation *FieldValidation `json:"validate,omitempty"`
}

func (field *Field) String() string {
	return "\t[" + field.SType + " " + field.Html + "]"
}

type SItems struct {
	Items []SItem
}

func (s *SItems) UnmarshalJSON(data []byte) (err error) {
	if data[0] == '{' {
		var items map[string]interface{}
		err = json.Unmarshal(data, &items)
		if err != nil {
			return
		}
		s.readSItems(0, items)
	} else {
		var list_of_items []map[string]interface{}
		err = json.Unmarshal(data, &list_of_items)
		if err != nil {
			return
		}

		for i, items := range list_of_items {
			s.readSItems(i, items)
		}
	}

	return
}

func (s *SItems) readSItems(id int, items map[string]interface{}) {
	var string_value string
	var i int = 0

	//append all items to list of SItems
	for key, value := range items {
		i++
		switch value.(type) {
		case string:
			string_value = value.(string)
		case int, int64:
			string_value = strconv.Itoa(value.(int))
		case float32:
			string_value = strconv.FormatFloat(value.(float64), 'f', 0, 32)
		case float64:
			string_value = strconv.FormatFloat(value.(float64), 'f', 0, 64)
		default:
			log.Printf("WARNING: field #%d: expecting a number or string, but JSON element #%d is %s", id+1, i, value)
			continue
		}

		sitem := &SItem{
			Key:   key,
			Value: string_value,
		}

		s.Items = append(s.Items, *sitem)
	}
}

type SItem struct {
	Key   string
	Value string
}
