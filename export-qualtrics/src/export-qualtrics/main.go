package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"text/template"
)

//arguments
var (
	input_file  string
	output_dir  string
	form_filter []string
	debug_mode  bool
	input_line  int = 0
)

var Templates template.Template

func main() {
	var err error

	fmt.Println("export-qualtrics v0.1 - convert Survana forms to Qualtrics TXT format")
	parse_arguments()

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s:%d: %s\n", input_file, input_line, r)
			debug.PrintStack()
			os.Exit(2)
		}
	}()

	db_file, err := os.Open(input_file)
	if err != nil {
		panic(err)
	}
	defer db_file.Close()

	reader := bufio.NewReader(db_file)

	//decode the JSON forms
	var (
		form         Form
		apply_filter bool = (len(form_filter) > 0)
		out          *os.File
	)

	for {
		input_line++

		//read one line from file
		form_string, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		//replace \uff0e with "." (hack. long story)
		form_string = strings.Replace(form_string, "．", ".", -1)

		//decode one form
		if err = json.Unmarshal([]byte(form_string), &form); err != nil {
			log.Printf("%s:%d: Error while decoding form: %s\n", input_file, input_line, err)
			continue
		}

		//apply filter
		if apply_filter && !has_form_id(form_filter, form.Id) {
			if debug_mode {
				log.Printf("%s:%d: Skipping form %s\n", input_file, input_line, form.Id)
			}
		}

		output_file := output_dir + "/" + form.Id + ".txt"

		//decide on the output (either use stdout, or open a file)
		if out, err = os.Create(output_file); err != nil {
			panic(err)
		}

		log.Printf("%s:%d: Converting form %s\n", input_file, input_line, form.Id)

		//convert the form
		if err = form.toQualtrics(out, &Templates); err != nil {
			out.Close()
			panic(err)
		}

		out.Close()
	}
}

func print_templates() {
	templates := Templates.Templates()
	num_templates := len(templates)

	var names string

	for i := 0; i < num_templates; i++ {
		names += templates[i].Name() + " "
	}

	log.Println("Known Survana templates: " + names)
}

func parse_arguments() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stdout, "ERROR: %s\n\n", r)
			show_usage()
			os.Exit(1)
		}
	}()
	var form_ids string
	flag.StringVar(&input_file, "i", "", "Database dump (required)")
	flag.StringVar(&output_dir, "o", "", "Output path")
	flag.StringVar(&form_ids, "f", "", "List of form IDs to convert (comma separated)")
	flag.BoolVar(&debug_mode, "g", false, "Debug mode (default: false)")
	flag.Parse()

	if len(input_file) == 0 {
		panic("Input file is required.")
	}

	if !file_exists(input_file) {
		panic(input_file + ": file not found")
	}

	if len(output_dir) == 0 {
		panic("Output folder path is required")
	}

	if !file_exists(output_dir) {
		err := os.MkdirAll(output_dir, 0771)
		if err != nil {
			panic(err)
		}
	}

	form_filter = strings.Split(form_ids, ",")
}

func file_exists(filepath string) bool {
	_, err := os.Stat(filepath)

	return err == nil
}

func has_form_id(forms []string, id string) bool {
	for _, form_id := range forms {
		if form_id == id {
			return true
		}
	}

	return false
}

func show_usage() {
	fmt.Println("Usage: " + os.Args[0] + " -i <FILE> [-o <DIR>] [-f ID1,ID2, ...]")
	flag.PrintDefaults()
}
