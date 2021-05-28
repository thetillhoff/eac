package apps

import (
	"fmt"
	"log"
	"runtime"
)

func Uninstall(appNames []string, shell string, appsDirPath string, continueOnError bool) {

	apps := apps(appNames, shell, continueOnError)

	for _, appItem := range apps {
		out, err := appItem.Uninstall(runtime.GOOS)
		fmt.Println(out)
		if err != nil {
			log.Fatal(err)
		}
	}
}
