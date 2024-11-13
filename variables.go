package main

import (
	"strings"
)

// ExtractVariables returns a list of unique variables found in a template string.
func ExtractVariables(templateStr string) []string {
	var variables []string

	for _, token := range strings.Fields(templateStr) {
		if strings.HasPrefix(token, "{{.") && strings.HasSuffix(token, "}}") {
			variable := strings.Trim(token, "{{.}}")
			if !contains(variables, variable) {
				variables = append(variables, variable)
			}
		}
	}

	return variables
}

func contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}
