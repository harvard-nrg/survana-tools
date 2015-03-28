package main

import (
	"fmt"
)

type FieldValidation struct {
	Required bool `json:"required"`
	Skip     bool `json:"skip"`
}

func (validation *FieldValidation) UnmarshalInterface(v interface{}) {
	validation_map := v.(map[string]interface{})
	unmarshalJSONBool(&validation.Required, validation_map["required"])
	unmarshalJSONBool(&validation.Skip, validation_map["skip"])
	if len(validation_map) > 2 {
		panic(fmt.Errorf("More than 2 items in s-validate object: %s", validation_map))
	}
}
