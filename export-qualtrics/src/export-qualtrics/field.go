package main

type Field struct {
    SId string `json:"s-id"`
    SType string `json:"s-type"`
    SItems map[string]int `json:"s-items"`
    SDirection string `json:"direction"`
    SGroup string `json:"s-group"`
    Html string `json:"html"`
    Theme string `json:"data-theme"`
    Validation *FieldValidation `json:"validate,omitempty"`
}

func (field *Field) String() string {
    return "\t[" + field.SType +" " + field.Html + "]"
}
