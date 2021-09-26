package templates

import (
	"log"
	"os"
)

func WriteTemplateToFile(filePath string, templateData interface{}, destinationPath string) {
	var (
		err     error
		f       *os.File
		content string
	)

	// Run templating
	content = RunTemplate(filePath, templateData)

	// Create file
	f, err = os.OpenFile(destinationPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0755)
	// f, err = os.Create(destinationPath)
	if err != nil {
		log.Fatalln("ERR Can't create file at '"+destinationPath+"';", err)
	}
	defer f.Close()

	// Write to file
	_, err = f.WriteString(content)
	if err != nil {
		log.Fatalln("ERR Can't write string to file at '"+destinationPath+"';", err)
	}

	// Ensure everything gets written
	f.Sync()
}
