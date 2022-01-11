package apps

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/thetillhoff/eac/internal/config"
	"github.com/thetillhoff/eac/pkg/apps/internal/app"
)

func downloadFile(url string, destination string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	out, err := os.Create(destination)
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

func downloadApp(conf config.Config, app app.App, platform string) {
	downloadFile("https://raw.githubusercontent.com/thetillhoff/eac/main/apps/"+app.Name+"/"+platform+"/getLatestVersion.sh", path.Join(AppsDirPath, app.Name, platform, "getLatestVersion.sh"))
}
