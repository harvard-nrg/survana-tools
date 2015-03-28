package main

import (
	"fmt"
	"text/template"
)

type qualtricsField struct {
	Question         *Field
	QuestionTemplate *template.Template
	QualtricsType    string
	Answer           *Field
	AnswerTemplate   *template.Template
}

const (
	MULTIPLE_CHOICE                 = "MC:SingleAnswer"
	MULTIPLE_CHOICE_MULTIPLE_ANSWER = "MC:MultipleAnswer"
	MULTIPLE_CHOICE_DROPDOWN        = "MC:DropDown"
	MATRIX                          = "Matrix"
	TE                              = "TE"
	TE_SINGLELINE                   = "TE:SingleLine"
	TE_MULTILINE                    = "TE:Essay"
	CS                              = "CS"
	RO                              = "RO"
	DB                              = "DB"
	SLIDER                          = "Slider"
)

func getQualtricsType(field *Field) (result string, err error) {
	switch field.SGroup {
	case "radio", "select":
		if field.SDirection == "horizontal" {
			return MULTIPLE_CHOICE + ":Horizontal", nil
		}
		return MULTIPLE_CHOICE, nil
	case "checkbox":
		if field.SDirection == "horizontal" {
			return MULTIPLE_CHOICE_MULTIPLE_ANSWER + ":Horizontal", nil
		}
		return MULTIPLE_CHOICE_MULTIPLE_ANSWER, nil
	}

	switch field.SType {
	case "input", "text", "number":
		return TE_SINGLELINE, nil
	case "textbox", "box":
		return TE_MULTILINE, nil
	case "question", "button", "label", "link", "html", "rtf":
		return DB, nil
	case "select", "toggle":
		return MULTIPLE_CHOICE_DROPDOWN, nil
	case "slider":
		return SLIDER, nil
	case "group":
		return MULTIPLE_CHOICE, nil
	}

	if len(field.SItems) > 0 {
		return MULTIPLE_CHOICE, nil
	}

	return "", fmt.Errorf("No Qualtrics answer type for field: %s", field)
}
