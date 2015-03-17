package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"
	"text/template"
)

//arguments
var (
	input_file       string
	output_file      string
	overwrite_output bool
	debug_mode       bool
)

var Templates template.Template

func main() {
	var err error

	fmt.Println("export-qualtrics v0.1 - convert Survana forms to Qualtrics TXT format")
	parse_arguments()

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", r)
			debug.PrintStack()
			os.Exit(2)
		}
	}()

	//read JSON form from file
	var input_data []byte
	if input_data, err = ioutil.ReadFile(input_file); err != nil {
		panic(err)
	}

	//decode the JSON form
	var form Form
	if err = json.Unmarshal(input_data, &form); err != nil {
		panic(err)
	}

	if debug_mode {
		print_templates()
		log.Println("\n" + form.String())
		return
	}

	//decide on the output (either use stdout, or open a file)
	var out *os.File
	if output_file == "-" {
		out = os.Stdout
	} else {
		if out, err = os.Create(output_file); err != nil {
			panic(err)
		}
		defer out.Close()
	}

	//convert the form
	if err = form.toQualtrics(out, &Templates); err != nil {
		panic(err)
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
	flag.StringVar(&input_file, "i", "", "Input file in Survana format (required)")
	flag.StringVar(&output_file, "o", "", "Output file in Qualtrics TXT format (default: add .qtxt to input file name)")
	flag.BoolVar(&overwrite_output, "f", false, "Overwrite output file if it exists (default: no)")
	flag.BoolVar(&debug_mode, "g", false, "Debug mode (default: false)")
	flag.Parse()

	if len(input_file) == 0 {
		panic("Input file is required.")
	}

	if len(output_file) == 0 {
		output_file = input_file + ".qtxt"
	}

	if !file_exists(input_file) {
		panic(input_file + ": file not found")
	}

	if !debug_mode && !overwrite_output && file_exists(output_file) {
		panic(output_file + ": file exists")
	}
}

func file_exists(filepath string) bool {
	_, err := os.Stat(filepath)

	return err == nil
}

func show_usage() {
	fmt.Println("Usage: " + os.Args[0] + " -i <FILE> [-o <FILE>] [-f]")
	flag.PrintDefaults()
}
