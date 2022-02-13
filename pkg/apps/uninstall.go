package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

// Calls the `uninstall` script of all the provided appnames
func Uninstall(appNames []string, appsDirPath string, verbose bool, versionsFilePath string) {
	logs.Verbose = verbose
	apps := parseApps(appNames, versionsFilePath)

	for _, appItem := range apps {
		logs.Info("Uninstalling app '" + appItem.Name + "'")
		out, err := appItem.Uninstall(appsDirPath, runtime.GOOS)
		if err == nil {
			logs.Success("Uninstalled app '" + appItem.Name + "'.")
		}
		logs.Info("Output of uninstallation script:", out)
		if err != nil {
			logs.Error("There was an error during uninstallation of app '"+appItem.Name+"':", err)
		}
	}
}
