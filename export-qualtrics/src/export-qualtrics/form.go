package main

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

type Form struct {
	Id           string  `json:"id"`
	Code         string  `json:"code"`
	Title        string  `json:"title"`
	Gid          string  `json:"gid"`
	GroupName    string  `json:"group"`
	CreatedOn    int64   `json:"created_on"`
	Version      string  `json:"version"`
	DisplayTitle bool    `json:"display_title"`
	Published    bool    `json:"published"`
	Fields       []Field `json:"data"`
}

func (form *Form) String() string {
	result := "[qualtrics form " + form.Id + "]\n"

	for i := 0; i < len(form.Fields); i++ {
		result += form.Fields[i].String() + "\n"
	}

	return result
}

func (form *Form) toQualtrics(out io.Writer, templates *template.Template) (err error) {
	out.Write([]byte("[[AdvancedFormat]]"))

	var (
		field, next_field *Field
		qualtrics_field   *qualtricsField
		qualtrics_type    string
		i                 int = 0
		num_fields            = len(form.Fields)
	)

	for i < num_fields {
		field = &form.Fields[i]

		//attempt to link a question with its answer, based on the assumption that most
		//answer fields follow a 'question' field. If a question is followed by a question,
		//render the first one as a simple instruction field
		if field.SType == "question" {
			if i < num_fields-1 {
				next_field = &form.Fields[i+1]
				if next_field.SType == "question" {
					qualtrics_field = &qualtricsField{
						Question:         field,
						QuestionTemplate: templates.Lookup("instruction"),
					}
				} else {
					//if this field has no html text, but answer has s-label, skip this field
					//and, on next iteration, convert the field label into a question
					if len(field.Html) == 0 && len(next_field.SLabel) > 0 {
						i++
						continue
					}
					qualtrics_type, err = getQualtricsType(next_field)
					if err != nil {
						return
					}
					qualtrics_field = &qualtricsField{
						Question:         field,
						QuestionTemplate: templates.Lookup("question"),
						QualtricsType:    qualtrics_type,
						Answer:           next_field,
						AnswerTemplate:   templates.Lookup(form.getTemplateName(next_field)),
					}
					i++
				}
			} else {
				qualtrics_field = &qualtricsField{
					Question:         field,
					QuestionTemplate: templates.Lookup("instruction"),
				}
			}
		} else {
			qualtrics_type, err = getQualtricsType(field)
			if err != nil {
				return
			}

			var question *Field = nil
			var question_tpl *template.Template = nil

			if len(field.SLabel) > 0 {
				question = &Field{
					SId:   field.SId,
					SType: "question",
					Html:  field.SLabel,
				}
				question_tpl = templates.Lookup("question")
			}

			qualtrics_field = &qualtricsField{
				Question:         question,
				QuestionTemplate: question_tpl,
				Answer:           field,
				QualtricsType:    qualtrics_type,
				AnswerTemplate:   templates.Lookup(form.getTemplateName(field)),
			}
		}

		//output question
		if qualtrics_field.Question != nil {
			if qualtrics_field.QuestionTemplate == nil {
				err = fmt.Errorf("Question template not found for field: %s", qualtrics_field.Question)
				return
			}
			err = qualtrics_field.QuestionTemplate.Execute(out, qualtrics_field)
			if err != nil {
				return
			}
		}

		//output answer
		if qualtrics_field.Answer != nil {
			if qualtrics_field.AnswerTemplate == nil {
				err = fmt.Errorf("Answer template not found for field: %s", qualtrics_field.Answer)
				return
			}
			err = qualtrics_field.AnswerTemplate.Execute(out, qualtrics_field)
			if err != nil {
				return
			}
		}

		i++
	}

	return
}

func (form *Form) getTemplateName(field *Field) string {
	if len(field.SType) > 0 {
		return field.SType
	} else if len(field.SGroup) > 0 {
		return field.SGroup + "_group"
	}

	log.Println("Failed to find template name for field: " + field.String())

	return ""
}
