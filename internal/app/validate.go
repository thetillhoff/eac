package app

import (
	"errors"
	"path"
	"strings"

	"github.com/thetillhoff/eac/pkg/logs"
)

func (app App) Validate(appsDirPath string, platform string) (string, error) {
	out := ""

	scriptpaths := []string{}
	scriptpaths = append(scriptpaths, path.Join(appsDirPath, app.Name, platform, app.getLocalVersionScript))
	scriptpaths = append(scriptpaths, path.Join(appsDirPath, app.Name, platform, app.installScript))
	scriptpaths = append(scriptpaths, path.Join(appsDirPath, app.Name, platform, app.configureScript))
	scriptpaths = append(scriptpaths, path.Join(appsDirPath, app.Name, platform, app.getLatestVersionScript))
	scriptpaths = append(scriptpaths, path.Join(appsDirPath, app.Name, platform, app.uninstallScript))

	missingFiles := testFiles(scriptpaths...)

	if len(missingFiles) == 0 {
		out = out + "The scripts for app '" + app.Name + "' exist.\n"
	} else {
		err := "The following files for app '" + app.Name + "' don't exist:\n"
		for _, missingFile := range missingFiles {
			err = err + missingFile + "\n"
		}
		return "", errors.New(err)
	}

	localVersion := app.LocalVersion(appsDirPath) // should the app not be installed locally, the this will return an empty string
	logs.Info("localVersion: ", localVersion)
	if strings.Contains(localVersion, "\n") {
		err := "The local version for app '" + app.Name + "' can't be retrieved. The result must be one line, but is:\n"
		err = err + localVersion
		return "", errors.New(err)
	} else if strings.Contains(localVersion, " ") {
		err := "The local version for app '" + app.Name + "' can't be retrieved. The result mustn't contain spaces, but is:\n"
		err = err + localVersion
		return "", errors.New(err)
	} else if localVersion == "" {
		err := "The local version for app '" + app.Name + "' couldn't be retrieved.\n"
		return "", errors.New(err)
	} else {
		out = out + "The local version for app '" + app.Name + "' could be retrieved.\n"
	}
	latestVersion := app.LatestVersion(appsDirPath, platform)
	logs.Info("latestVersion: ", latestVersion)
	if strings.Contains(latestVersion, "\n") {
		err := "The latest version for app '" + app.Name + "' can't be retrieved. The result should be one line, but is:\n"
		err = err + latestVersion
		return "", errors.New(err)
	} else {
		out = out + "The latest version for app '" + app.Name + "' could be retrieved.\n"
	}

	return out, nil
}
