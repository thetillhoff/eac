package app

import (
	"errors"
	"path"
	"strings"
)

func (app App) Validate(appsDirPath string, platform string) (string, error) {
	out := ""

	shellpath := []string{app.shell}
	if strings.Contains(app.shell, " ") { // if shell contains spaces (and therefore arguments)
		// don't edit the app.shell directly, as it is later used to retrieve the local and latest version. Instead edit the copy in shellpath.
		shellpath = []string{strings.Split(app.shell, " ")[0]} // remove those arguments (before checking for file existance)
	}

	if len(testFiles(shellpath...)) == 0 {
		out = out + "The shell for '" + app.Name + "' exists.\n"
	} else {
		err := "The shell '" + app.shell + "' for app '" + app.Name + "' doesn't exist."
		return "", errors.New(err)
	}

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
	if strings.Contains(localVersion, "\n") {
		err := "The local version for app '" + app.Name + "' can't be retrieved. The result should be one line, but is:\n"
		err = err + localVersion
		return "", errors.New(err)
	} else {
		out = out + "The local version for app '" + app.Name + "' could be retrieved.\n"
	}
	latestVersion := app.LatestVersion(appsDirPath, platform)
	if strings.Contains(latestVersion, "\n") {
		err := "The latest version for app '" + app.Name + "' can't be retrieved. The result should be one line, but is:\n"
		err = err + latestVersion
		return "", errors.New(err)
	} else {
		out = out + "The latest version for app '" + app.Name + "' could be retrieved.\n"
	}

	return out, nil
}
