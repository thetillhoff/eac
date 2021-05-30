package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

func Uninstall(appNames []string, shell string, appsDirPath string, continueOnError bool) {

	apps := apps(appNames, shell, continueOnError)

	for _, appItem := range apps {
		out, err := appItem.Uninstall(appsDirPath, runtime.GOOS)
		logs.Info("Output of configuration script:", out)
		if err != nil {
			logs.Err("There was an error during uninstallation of app '"+appItem.Name+"':", err)
		}
	}
}
