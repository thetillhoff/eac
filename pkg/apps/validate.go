package apps

import (
	"fmt"
	"log"
)

func Validate(appNames []string, flaggedPlatforms []string, shell string, appsDirPath string, continueOnError bool) {

	apps := apps(appNames, shell, continueOnError)

	platforms := ResolvePlatforms(flaggedPlatforms)

	for _, appItem := range apps {
		for _, platform := range platforms {
			out, err := appItem.Validate(appsDirPath, platform)
			fmt.Println(out)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
