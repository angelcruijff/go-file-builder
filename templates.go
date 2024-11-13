package main

import (
	"fmt"
	"os"
	"text/template"
)

// ReadTemplateFile reads the content of a template file.
func ReadTemplateFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ExecuteTemplate(templateContent string, data interface{}, newFile string) error {
	print_file := false
	tmpl, err := template.New("myTemplate").Parse(templateContent)
	if err != nil {
		return err
	}

	if newFile == "" {
		//If the new_file parameter is not given, output the final file to the OS Output
		print_file = true
	}

	if print_file {
		err = tmpl.Execute(os.Stdout, data)
	} else {
		created_file, err := os.Create(newFile)
		if err != nil {
			panic(err)
		}
		defer created_file.Close()
		err = tmpl.Execute(created_file, data)
		return err
	}

	return err // Convert the buffer to a string
}

func createDataMap(variables []string, strictMode bool) map[string]string {
	data := make(map[string]string)

	for _, variable := range variables {
		value, varPresent := os.LookupEnv(variable)
		if varPresent {
			data[variable] = value
		} else if strictMode {
			handleError(fmt.Errorf("The variable '%s' is not set", variable))
		}
	}

	return data
}
