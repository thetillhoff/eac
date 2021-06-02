package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

func Install(appNames []string, noConfigure bool, update bool, shell string, appsDirPath string, continueOnError bool, versionsFilePath string, latest bool, verbose bool) {
	logs.Verbose = verbose
	apps := apps(appNames, shell, continueOnError)

	// if not latest
	//   if not update
	//     loadVersions from file
	//   else if update
	//     updateAppVersion (and write back to )
	// for each app
	//   if latest
	//     pass "" as version to appItem.Install to get latest version
	//   if not latest
	//     pass version either from versionsFile or from app.wantedVersion (latter has prio) // if app.wantedVersion is set, and latest flag is set, warn about ignoring "latest" arg (this allows to install some apps with version and some with latest in the same command)

	if !latest { // if specific version should be installed (or version was never specified)
		loadVersions(versionsFilePath, continueOnError)
		if update {
			for _, appItem := range apps {
				err := updateAppVersion(appItem, appsDirPath, runtime.GOOS, continueOnError, versionsFilePath, false, true) // Update app version, don't care about currently installed version
				if err != nil {
					logs.Err("There was an error while updating the version for app '"+appItem.Name+"':", continueOnError, err)
				}
			}
		}
	}

	for _, appItem := range apps {
		if appItem.LocalVersion(appsDirPath) != "" && appItem.LocalVersion(appsDirPath) == appItem.WantedVersion() {
			logs.Info("App '" + appItem.Name + "' is already installed in wanted version '" + appItem.WantedVersion() + "'.")
		} else { // app is either not installed, or installed in wrong version
			appItem, out, err := appItem.Install(appsDirPath, runtime.GOOS, "") // Install app

			if err == nil {
				logs.Success("Installed app '" + appItem.Name + "' in version '" + appItem.WantedVersion() + "'.")
			}
			if out != "" {
				logs.Info("Output of installation script:", out)
			}
			if err != nil {
				logs.Err("There was an error while installing the app '"+appItem.Name+"':", continueOnError, err)
			}

			if !noConfigure {
				out, err := appItem.Configure(appsDirPath, runtime.GOOS) // Configure app
				if err == nil {
					logs.Success("Configured app '" + appItem.Name + "'.")
				}
				if out != "" {
					logs.Info("Output of configuration script:", out)
				}
				if err != nil {
					logs.Err("There was an error during configuration of app '"+appItem.Name+"':", continueOnError, err)
				}
			}
		}
	}
}
