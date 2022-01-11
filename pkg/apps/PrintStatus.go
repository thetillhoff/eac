package apps

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/thetillhoff/eac/internal/config"
)

var (
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
)

func PrintStatus(conf config.Config) {
	apps, _ := List(conf)

	for _, app := range apps {
		fmt.Print(app.Name)
		if app.InstalledVersion(conf.AppsDirPath) == "" {
			// app is not installed
		} else if app.InstalledVersion(conf.AppsDirPath) == app.LatestVersion(conf.AppsDirPath, runtime.GOOS) {
			fmt.Print("=")
			green.Fprint(os.Stdout, app.InstalledVersion(conf.AppsDirPath))
		} else {
			fmt.Print("=")
			yellow.Fprint(os.Stdout, app.InstalledVersion(conf.AppsDirPath))
		}
		fmt.Print("\n")
	}
}
