package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Parse command-line flags
	fileName, newFile, strictMode := parseFlags()

	// Read the template file
	templateContent, err := readTemplateFile(fileName)
	if err != nil {
		handleError(err)
	}

	// Create a new template
	tmpl, err := template.New("myTemplate").Parse(templateContent)
	if err != nil {
		handleError(err)
	}

	// Extract variables from the template
	variables := extractVariables(templateContent)

	// Create a data map with values from environment variables
	data := createDataMap(variables, strictMode)

	// Execute the template
	err = executeTemplate(tmpl, data, newFile)
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
