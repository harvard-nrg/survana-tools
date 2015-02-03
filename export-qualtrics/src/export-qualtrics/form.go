package main

import (
        "io"
        "log"
        "text/template"
       )

type Form struct {
    Id string `json:"id"`
    Code string `json:"code"`
    Title string `json:"title"`
    Gid string `json:"gid"`
    GroupName string `json:"group"`
    CreatedOn int64 `json:"created_on"`
    Version string `json:"version"`
    DisplayTitle bool `json:"display_title"`
    Published bool `json:"published"`
    Fields []Field `json:"data"`
}

func (form *Form) String() string {
    result := "[qualtrics form " + form.Id + "]\n"

    for i := 0; i < len(form.Fields); i++ {
        result += form.Fields[i].String() + "\n"
    }
    
    return result
}

func (form *Form) toQualtrics(out io.Writer, templates *template.Template) (err error) {

    for i := 0; i < len(form.Fields); i++ {
        field := form.Fields[i]
        tpl := templates.Lookup(field.SType)
        if (tpl == nil) {
            log.Println("template '" + field.SType + "' not found")
            continue
        }
        
        err = tpl.Execute(out, field)
        if err != nil {
            return
        }
    }
    

    return
}
