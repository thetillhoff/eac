package apps

import (
	"fmt"
	"os"
	"path"
)

func Clean(appNames []string) {
	if len(appNames) > 0 { // Remove specific apps only
		for _, appName := range appNames {
			appDir := path.Join(workdir, appName)
			if Verbose {
				fmt.Println("Removing all cached versions for app", appName, "at", appDir, "...")
			}

			if !DryRun { // Not if dry-running
				os.RemoveAll(appDir) // Remove /tmp/eac
			}

			if Verbose {
				fmt.Println("Removal finished.")
			}
		}
	} else { // Remove whole workdir by default
		if Verbose {
			fmt.Println("Removing workdir at", workdir, "...")
		}

		if !DryRun { // Not if dry-running
			os.RemoveAll(workdir) // Remove /tmp/eac
		}

		if Verbose {
			fmt.Println("Removal finished.")
		}
	}
}
