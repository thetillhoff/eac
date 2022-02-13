package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

// Calls the `configure` script of all the provided appnames
func Configure(appNames []string, appsDirPath string, verbose bool, checkInstalledVersion bool, versionsFilePath string) {
	logs.Verbose = verbose
	apps := parseApps(appNames, versionsFilePath)

	for _, appItem := range apps {
		logs.Info("Configuring app '" + appItem.Name + "'")
		if checkInstalledVersion {
			installedVersion := appItem.InstalledVersion(appsDirPath) // test if get-local-version works (to check if app is already installed), if not, fail (with custom error)
			if installedVersion == "" {
				logs.Error("There was an error during configuration of app '"+appItem.Name+"':", "It was not possible to retrieve the local version.\nAre you sure the scripts are up-to-date?")
			}
		}

		out, err := appItem.Configure(appsDirPath, runtime.GOOS) // Configure app
		if err == nil {
			logs.Success("Configured app '" + appItem.Name + "'.")
		}
		if out != "" {
			logs.Info("Output of configuration script:", out)
		}
		if err != nil {
			logs.Error("There was an error during configuration of app '"+appItem.Name+"':", err)
		}
	}
}
