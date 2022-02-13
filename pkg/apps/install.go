package apps

import (
	"os"
	"path"
	"runtime"

	"github.com/thetillhoff/eac/internal/config"
	"github.com/thetillhoff/eac/pkg/apps/internal/appVersions"
	"github.com/thetillhoff/eac/pkg/logs"
)

// Calls the `install` script of all the provided appnames
func Install(appNames []string, conf config.Config) {
	logs.Verbose = conf.Verbose
	apps := parseApps(appNames, conf.VersionsFilePath)

	// check if folder for app exist
	//   if not, ask whether the folder should be created and they should be downloaded
	//   else skip this app

	for _, appItem := range apps { // for each app

		if _, err := os.Stat(path.Join(conf.AppsDirPath, appItem.Name)); os.IsNotExist(err) { // Check if folder for app exists
			// create the folder and download the files
			Create([]string{appItem.Name}, []string{runtime.GOOS}, conf.AppsDirPath, conf.Verbose, map[string]string{})
			downloadApp(conf, appItem, runtime.GOOS)
		} // Else folder for app does not exist, then the app is available locally -> nothing to do here

		if conf.Latest { // If latest version should be installed
			if conf.Update {
				appItem = appVersions.Update(appItem, conf.AppsDirPath, conf.VersionsFilePath) // Update app version, don't care about currently installed version
			}
			appItem.WantedVersion = appItem.LatestVersion(conf.AppsDirPath, runtime.GOOS) // Set WantedVersion to the latest available one
			appVersions.SetVersion(appItem.Name, appItem.WantedVersion)
			appVersions.Save(conf.VersionsFilePath)
			logs.Info("Due to flag 'latest', the wantedVersion of app '" + appItem.Name + "' was set to 'v" + appItem.WantedVersion + "'.")
		} else if appItem.WantedVersion == "" { // If no version was specified
			appItem.WantedVersion = appItem.LatestVersion(conf.AppsDirPath, runtime.GOOS) // Set WantedVersion to the latest available one
			appVersions.SetVersion(appItem.Name, appItem.WantedVersion)
			appVersions.Save(conf.VersionsFilePath)
			logs.Info("Version for app '" + appItem.Name + "' was automatically set to latest 'v" + appItem.WantedVersion + "'.")
		}
		// implicit else:
		//   pass version already existant in app.WantedVersion (f.e. from versionsFile)

		if appItem.InstalledVersion(conf.AppsDirPath) == appItem.WantedVersion { // If app is already installed in desired version
			logs.Success("App '" + appItem.Name + "' is already installed in wanted version '" + appItem.WantedVersion + "'.")
		} else { // app is either not installed, or installed in wrong version
			appItem, out, err := appItem.Install(conf.AppsDirPath, runtime.GOOS, "") // Install app

			if err == nil {
				logs.Success("Installed app '" + appItem.Name + "' in version '" + appItem.WantedVersion + "'.")
			}
			if out != "" {
				logs.Info("Output of installation script:", out)
			}
			if err != nil {
				logs.Error("There was an error while installing the app '"+appItem.Name+"':", err)
			}

			if !conf.NoConfiguration {
				out, err := appItem.Configure(conf.AppsDirPath, runtime.GOOS) // Configure app
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
	}
}
