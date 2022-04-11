package eac

import "github.com/thetillhoff/eac/pkg/eac/internal/apps"

func ListAvailableApps() ([]string, error) {
	appNameList := []string{}

	for appName := range apps.Apps {
		appNameList = append(appNameList, appName)
	}

	return appNameList, nil
}
