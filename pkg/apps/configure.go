package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

func Configure(appNames []string, appsDirPath string, verbose bool, checkLocalVersion bool, versionsFilePath string) {
	logs.Verbose = verbose
	apps := apps(appNames, versionsFilePath)

	for _, appItem := range apps {
		if checkLocalVersion {
			localVersion := appItem.LocalVersion(appsDirPath) // test if get-local-version works (to check if app is already installed), if not, fail (with custom error)
			if localVersion == "" {
				logs.Err("There was an error during configuration of app '"+appItem.Name+"':", "It was not possible to retrieve the local version.\nAre you sure the scripts are up-to-date?")
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
			logs.Err("There was an error during configuration of app '"+appItem.Name+"':", err)
		}
	}
}
