package main

import (
        "io"
        "text/template"
       )


type QTXTMarshaller interface {
    toQTXT(io.Writer, *template.Template) error
}
