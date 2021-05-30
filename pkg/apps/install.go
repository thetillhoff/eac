package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

func Install(appNames []string, noConfigure bool, update bool, shell string, appsDirPath string, continueOnError bool, versionsFilePath string, latest bool) {
	apps := apps(appNames, shell, continueOnError)

	if !latest { // if specific version should be installed (or version was never specified)
		loadVersions(versionsFilePath)
		if update {
			for _, appItem := range apps {
				err := updateAppVersion(appItem, appsDirPath, runtime.GOOS, versionsFilePath, false, true) // Update app version, don't care about currently installed version
				if err != nil {
					logs.Err("There was an error while updating the version for app '"+appItem.Name+"':", err)
				}
			}
		}
	}

	for _, appItem := range apps {
		if appItem.LocalVersion() != "" && appItem.LocalVersion() == appItem.WantedVersion() {
			logs.Info("App '" + appItem.Name + "' is already installed in wanted version '" + appItem.WantedVersion() + "'.")
		} else { // app is either not installed, or installed in wrong version
			out, err := appItem.Install(appsDirPath, runtime.GOOS, latest) // Install app
			logs.Info("Output of configuration script:", out)
			if err != nil {
				logs.Err("There was an error while installing the app '"+appItem.Name+"':", err)
			}

			if !noConfigure {
				out, err := appItem.Configure(appsDirPath, runtime.GOOS) // Configure app
				logs.Info("Output of configuration script:", out)
				if err != nil {
					logs.Err("There was an error while configuring the app '"+appItem.Name+"':", err)
				}
			}
		}
	}
}
