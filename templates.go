// template.go
package main

import (
	"bytes"
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

func ExecuteTemplate(templateContent string, data interface{}) (string, error) {
	tmpl, err := template.New("myTemplate").Parse(templateContent)
	if err != nil {
		return "", err
	}

	var outputBuffer bytes.Buffer // Use a buffer to capture the output
	err = tmpl.Execute(&outputBuffer, data)
	if err != nil {
		return "", err
	}

	return outputBuffer.String(), nil // Convert the buffer to a string
}
