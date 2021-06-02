package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

func Uninstall(appNames []string, shell string, appsDirPath string, continueOnError bool, verbose bool) {
	logs.Verbose = verbose
	apps := apps(appNames, shell, continueOnError)

	for _, appItem := range apps {
		out, err := appItem.Uninstall(appsDirPath, runtime.GOOS)
		if err == nil {
			logs.Success("Uninstalled app '" + appItem.Name + "'.")
		}
		logs.Info("Output of uninstallation script:", out)
		if err != nil {
			logs.Err("There was an error during uninstallation of app '"+appItem.Name+"':", continueOnError, err)
		}
	}
}
