package templates

import (
	"io/fs"
	"log"
	"os"
	"strings"
)

func GetPlatforms() []string {
	var (
		err       error
		files     fs.FS
		platforms []string
	)

	files, err = fs.Sub(templates, "files")
	if err != nil {
		log.Fatalln("ERR Can't remove prefix from template names;", err)
	}

	fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		newPlatform := strings.Split(path, string(os.PathSeparator))[0]
		exists := false
		for _, platform := range platforms {
			if platform == newPlatform {
				exists = true
			}
		}
		if !exists && newPlatform != "." {
			platforms = append(platforms, newPlatform)
		}
		return err
	})

	return platforms
}
