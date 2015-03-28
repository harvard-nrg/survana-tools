package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func unmarshalJSONString(s *string, v interface{}) {
	switch v.(type) {
	case nil:
		*s = ""
	case string:
		*s = v.(string)
	case int:
		*s = strconv.FormatInt(int64(v.(int)), 10)
	case int32:
		*s = strconv.FormatInt(int64(v.(int32)), 10)
	case int64:
		*s = strconv.FormatInt(v.(int64), 10)
	case float32:
		*s = strconv.FormatInt(int64(v.(float32)), 10)
	case float64:
		*s = strconv.FormatInt(int64(v.(float64)), 10)
	case bool:
		*s = strconv.FormatBool(v.(bool))
	default:
		panic(fmt.Errorf("Don't know how to unmarshal [string] value: %s", v))
	}
}

func unmarshalJSONInt64(i *int64, v interface{}) {
	switch v.(type) {
	case nil:
		*i = 0
	case int64:
		*i = v.(int64)
	case int32:
		*i = int64(v.(int32))
	case int:
		*i = int64(v.(int))
	case float64:
		*i = int64(v.(float64))
	case float32:
		*i = int64(v.(float32))
	case string:
		i32, err := strconv.Atoi(v.(string))
		if err != nil {
			panic(fmt.Errorf("Cannot convert string '%s' to int64: %s", v.(string), err))
		}
		*i = int64(i32)
	case bool:
		if v.(bool) {
			*i = 1
		} else {
			*i = 0
		}
	default:
		panic(fmt.Errorf("Don't know how to unmarshal [int64] value: %s", v))
	}
}

func unmarshalJSONBool(b *bool, v interface{}) {
	switch v.(type) {
	case nil:
		*b = false
	case bool:
		*b = v.(bool)
	case int64:
		if v.(int64) == 0 {
			*b = false
		} else {
			*b = true
		}
	case int32:
		if v.(int32) == 0 {
			*b = false
		} else {
			*b = true
		}
	case int:
		if v.(int) == 0 {
			*b = false
		} else {
			*b = true
		}
	case float64:
		if v.(float64) < 0.0001 {
			*b = false
		} else {
			*b = true
		}
	case float32:
		if v.(float32) < 0.0001 {
			*b = false
		} else {
			*b = true
		}
	case string:
		lc := strings.TrimSpace(strings.ToLower(v.(string)))

		switch lc {
		case "true", "yes", "y":
			*b = true
		default:
			*b = false
		}

	default:
		panic(fmt.Errorf("Don't know how to unmarshal [bool] value: %s", v))
	}
}

func unmarshalJSONFields(fields *[]Field, v interface{}) {
	if v == nil {
		return
	}
	bytes, _ := json.Marshal(v)
	j := string(bytes)
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		if debug_mode {
			log.Printf("unmarshalling slice containing %d Fields: %s", s.Len(), j)
		}
		for i := 0; i < s.Len(); i++ {
			if debug_mode {
				log.Printf("unmarshalling slice %d of %d\n", i+1, s.Len())
			}
			unmarshalJSONFields(fields, s.Index(i).Interface())
		}
	case reflect.Map:
		field_map := v.(map[string]interface{})

		_, has_id := field_map["s-id"]
		_, has_type := field_map["s-type"]
		_, has_items := field_map["s-items"]
		_, has_group := field_map["s-group"]

		//{"s-id":x, "s-items":[...] }
		if has_id || has_type || has_items || has_group {
			if debug_mode {
				log.Printf("Unmarshalling field: %s\n", j)
			}
			var field Field
			field.UnmarshalInterface(v)
			*fields = append(*fields, field)
		} else {
			if debug_mode {
				log.Printf("Unmarshalling key-value pairs: %s\n", j)
			}
			//{"key":value, "key2":value, etc}
			for k, v := range field_map {
				var field Field
				field.UnmarshalKeyValueInterface(k, v)
				*fields = append(*fields, field)
			}

			//sort fields by value
			sort.Sort(ByFieldValue(*fields))
		}

	default:
		panic(fmt.Errorf("Don't know how to unmarshal [Field] value: %s", j))
	}
}

func unmarshalJSONFieldValidation(validation **FieldValidation, v interface{}) {
	if v == nil {
		return
	}

	if reflect.TypeOf(v).Kind() != reflect.Map {
		panic(fmt.Errorf("Field validation is not a map!"))
	}

	if *validation == nil {
		*validation = &FieldValidation{}
	}

	(*validation).UnmarshalInterface(v)
}
