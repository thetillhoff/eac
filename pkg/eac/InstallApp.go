package apps

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/thetillhoff/eac/pkg/eac/internal/apps"
	"github.com/thetillhoff/eac/pkg/goCopy"
)

func InstallApp(appName string, version string) error {
	var (
		err        error
		sourcePath string
		app        = apps.Apps[appName]
	)

	fmt.Println("Installing", appName+"@"+version)

	sourcePath, err = DownloadAppVersion(appName, version)
	if err != nil {
		os.RemoveAll(path.Join(workdir, appName, version)) // Remove version folder on error
		return err
	}
	if Verbose {
		fmt.Println("Download finished.")
	}

	if _, err = os.Stat(app.Destination); err == nil { // Check if destination exists
		// destination exists
		if Verbose {
			fmt.Println("Removing old version of " + appName + " ...")
		}

		if !DryRun { // Not if dry-running
			err = os.RemoveAll(app.Destination)
			if err != nil {
				return err
			}
		}

		if Verbose {
			fmt.Println("Removal finished.")
		}

	} else if errors.Is(err, os.ErrNotExist) {
		// destination does not exist -> nothing to do
	} else { // Schroedinger
		return errors.New("could not access destination: " + app.Destination)
	}

	if Verbose {
		fmt.Println("Placing " + appName + "@" + version + " ...")
	}

	if !DryRun { // Not if dry-running
		goCopy.Verbose = Verbose
		err = goCopy.Copy(sourcePath, app.Destination)
		if err != nil {
			return err
		}
	}

	if Verbose {
		fmt.Println("Placement finished.")
	}

	if len(app.PermissionAdjustments) > 0 {
		if Verbose {
			fmt.Println("Setting permissions...")
		}

		for _, filePermission := range app.PermissionAdjustments {
			destinationPath := path.Join(path.Dir(app.Destination), filePermission.Path)
			if !DryRun { // Not if dry-running
				err = os.Chmod(destinationPath, filePermission.Mode)
				if err != nil {
					return errors.New("could not set destination permissions: " + err.Error())
				}
			}
		}

		if Verbose {
			fmt.Println("Permissions set.")
		}
	}

	return nil
}
