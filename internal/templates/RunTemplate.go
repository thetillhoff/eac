package templates

import (
	"bytes"
	"io/fs"
	"log"
	"text/template"
)

func RunTemplate(filePath string, templateData interface{}) string {
	var (
		err    error
		files  fs.FS
		buffer bytes.Buffer
	)

	files, err = fs.Sub(templates, "files")
	if err != nil {
		log.Fatalln("ERR Can't remove prefix from template names;", err)
	}

	t, err := template.ParseFS(files, filePath)
	if err != nil {
		log.Fatalln("ERR Couldn't load template at '"+filePath+"';", err)
	}
	t.Execute(&buffer, templateData)

	return buffer.String()
}
