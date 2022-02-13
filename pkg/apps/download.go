package apps

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/thetillhoff/eac/internal/config"
	"github.com/thetillhoff/eac/pkg/apps/internal/app"
	"github.com/thetillhoff/eac/pkg/logs"
)

func downloadFile(url string, destination string) {

	logs.Info("url:", url)
	logs.Info("destination:", destination)

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
	logs.Info("Downloading files for app '" + app.Name + "'")
	downloadFile("https://raw.githubusercontent.com/thetillhoff/eac/main/apps/"+app.Name+"/"+platform+"/getLatestVersion.sh", path.Join(conf.AppsDirPath, app.Name, platform, "getLatestVersion.sh"))
	downloadFile("https://raw.githubusercontent.com/thetillhoff/eac/main/apps/"+app.Name+"/"+platform+"/getLocalVersion.sh", path.Join(conf.AppsDirPath, app.Name, platform, "getLocalVersion.sh"))
	downloadFile("https://raw.githubusercontent.com/thetillhoff/eac/main/apps/"+app.Name+"/"+platform+"/install.sh", path.Join(conf.AppsDirPath, app.Name, platform, "install.sh"))
	downloadFile("https://raw.githubusercontent.com/thetillhoff/eac/main/apps/"+app.Name+"/"+platform+"/configure.sh", path.Join(conf.AppsDirPath, app.Name, platform, "configure.sh"))
	downloadFile("https://raw.githubusercontent.com/thetillhoff/eac/main/apps/"+app.Name+"/"+platform+"/uninstall.sh", path.Join(conf.AppsDirPath, app.Name, platform, "uninstall.sh"))
}
