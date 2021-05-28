package apps

import (
	"log"
	"runtime"
)

func Update(appNames []string, shell string, appsDirPath string, continueOnError bool, versionsFilePath string, dryRun bool, skipLocal bool) {

	apps := apps(appNames, shell, continueOnError)

	for _, appItem := range apps {
		err := updateAppVersion(appItem, runtime.GOOS, versionsFilePath, dryRun, skipLocal)
		if err != nil {
			log.Fatalln("There was an error while updating the version for app '" + appItem.Name + "':\n" + err.Error())
		}
	}
}
