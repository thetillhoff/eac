package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/internal/appVersions"
	"github.com/thetillhoff/eac/internal/config"
	"github.com/thetillhoff/eac/pkg/logs"
)

func Install(appNames []string, conf config.Config) {
	logs.Verbose = conf.Verbose
	apps := apps(appNames, conf.VersionsFilePath)

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

	if !conf.Latest { // if specific version should be installed (or version was never specified)
		if conf.Update {
			for _, appItem := range apps {
				appItem = appVersions.Update(appItem) // Update app version, don't care about currently installed version
			}
		}
	}

	for _, appItem := range apps {
		if conf.Latest {
			appItem.WantedVersion = appItem.LatestVersion(conf.AppsDirPath, runtime.GOOS)
			appVersions.SetVersion(appItem.Name, appItem.WantedVersion)
			appVersions.Save(conf.VersionsFilePath)
			logs.Info("Due to flag 'latest', the wantedVersion of app '" + appItem.Name + "' was set to 'v" + appItem.WantedVersion + "'.")
		} else if appItem.WantedVersion == "" {
			appItem.WantedVersion = appItem.LatestVersion(conf.AppsDirPath, runtime.GOOS)
			appVersions.SetVersion(appItem.Name, appItem.WantedVersion)
			appVersions.Save(conf.VersionsFilePath)
			logs.Info("Version for app '" + appItem.Name + "' was automatically set to latest 'v" + appItem.WantedVersion + "'.")
		}
		if appItem.LocalVersion(conf.AppsDirPath) == appItem.WantedVersion {
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
				logs.Err("There was an error while installing the app '"+appItem.Name+"':", err)
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
					logs.Err("There was an error during configuration of app '"+appItem.Name+"':", err)
				}
			}
		}
	}
}
