package templates

import (
	"io/fs"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/thetillhoff/eac/pkg/logs"
)

func GetSharedFiles() []string {
	var (
		err         error
		files       fs.FS
		sharedFiles []string
	)

	files, err = fs.Sub(templates, "files")
	if err != nil {
		log.Fatalln("ERR Can't remove prefix from template names;", err)
	}

	fs.WalkDir(files, ".", func(inputPath string, d fs.DirEntry, err error) error {
		if inputPath != "." && strings.Split(inputPath, string(os.PathSeparator))[0] == runtime.GOOS && strings.HasPrefix(path.Base(inputPath), "shared-") {
			logs.Info("Detected shared file with name '" + path.Base(inputPath) + "'.")
			sharedFiles = append(sharedFiles, path.Base(inputPath))
		}
		return err
	})

	return sharedFiles
}
