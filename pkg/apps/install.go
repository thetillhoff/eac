package apps

import (
	"fmt"
	"log"
	"runtime"
)

func Install(appNames []string, noConfigure bool, update bool, shell string, appsDirPath string, continueOnError bool, versionsFilePath string, latest bool) {
	apps := apps(appNames, shell, continueOnError)

	if !latest { // if specific version should be installed (or version was never specified)
		loadVersions(versionsFilePath)
		if update {
			for _, appItem := range apps {
				err := updateAppVersion(appItem, runtime.GOOS, versionsFilePath, false, true) // Update app version, don't care about currently installed version
				if err != nil {
					log.Fatalln("There was an error while updating the version for app '" + appItem.Name + "':\n" + err.Error())
				}
			}
		}
	}

	for _, appItem := range apps {
		if appItem.LocalVersion() != "" && appItem.LocalVersion() == appItem.WantedVersion() {
			fmt.Println("App '" + appItem.Name + "' is already installed in wanted version '" + appItem.WantedVersion() + "'.")
		} else { // app is either not installed, or installed in wrong version
			out, err := appItem.Install(runtime.GOOS, latest) // Install app
			fmt.Println(out)
			if err != nil {
				log.Fatalln("There was an error while installing the app '" + appItem.Name + "':\n" + err.Error())
			}

			if !noConfigure {
				out, err := appItem.Configure(runtime.GOOS) // Configure app
				fmt.Println(out)
				if err != nil {
					log.Fatal("There was an error while configuring the app '" + appItem.Name + "':\n" + err.Error())
				}
			}
		}
	}
}
