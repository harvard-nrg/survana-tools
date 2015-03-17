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
	MULTIPLE_CHOICE                 = "MC"
	MULTIPLE_CHOICE_MULTIPLE_ANSWER = "MC:MultipleAnswer"
	MATRIX                          = "Matrix"
	TE                              = "TE"
	TE_SINGLELINE                   = "TE:SingleLine"
	TE_MULTILINE                    = "TE:Essay"
	CS                              = "CS"
	RO                              = "RO"
	DB                              = "DB"
)

func getQualtricsType(field *Field) (result string, err error) {
	switch field.SType {
	case "input", "text", "number":
		return TE_SINGLELINE, nil
	case "textbox":
		return TE_MULTILINE, nil
	}

	switch field.SGroup {
	case "radio", "select":
		return MULTIPLE_CHOICE, nil
	case "checkbox":
		return MULTIPLE_CHOICE_MULTIPLE_ANSWER, nil
	}

	return "", fmt.Errorf("No Qualtrics answer type for field: %s", field)
}
