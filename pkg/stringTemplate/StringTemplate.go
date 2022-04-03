package stringTemplate

import (
	"bytes"
	"text/template"
)

func StringTemplate(source string, mappedValues map[string]interface{}) (string, error) {

	tmpl, err := template.New("template").Parse(source) //  .Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		return "", err
	}

	outputBuffer := new(bytes.Buffer)
	outputBuffer.Reset()
	err = tmpl.Execute(outputBuffer, mappedValues)
	if err != nil {
		return "", err
	}

	return outputBuffer.String(), nil
}
