package eac

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/thetillhoff/eac/pkg/eac/internal/apps"
)

func ListDownloadedApps(withVersions bool) ([]string, error) {
	var (
		err           error
		appName       string
		appNameList   = []string{}
		appPath       string
		appFolder     []fs.FileInfo
		versionFolder fs.FileInfo
	)

	for appName = range apps.Apps {
		appPath = path.Join(workdir, appName)

		if _, err = os.Stat(appPath); err == nil { // Check if appFolder exists
			if !withVersions { // Only apps, no versions

				appNameList = append(appNameList, appName)

			} else { // Apps with versions

				appFolder, err = ioutil.ReadDir(path.Join(workdir, appName)) //
				if err != nil {
					return appNameList, err
				}

				for _, versionFolder = range appFolder {
					appNameList = append(appNameList, appName+"@"+versionFolder.Name())
				}
			}
		}
	}

	return appNameList, nil
}
