package main

import (
	"fmt"
	"os"
)

func main() {
	// Parse command-line flags
	flags := ParseFlags()
	// Access individual flag values
	fileName := flags.FileName
	newFile := flags.NewFileName
	strictMode := flags.StrictMode

	// Read the template file
	templateContent, err := ReadTemplateFile(fileName)
	if err != nil {
		handleError(err)
	}

	// Extract variables from the template
	variables := ExtractVariables(templateContent)

	// Create a data map with values from environment variables
	data := createDataMap(variables, strictMode)

	// Execute the template
	err = ExecuteTemplate(templateContent, data, newFile)
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
