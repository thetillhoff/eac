package eac

import (
	"fmt"
	"os"
	"path"
)

// Removes either the whole workdir (/tmp/eac) in case no args are provided,
// or removes only contents of it in case args are provided. The args reflect the name of the entry that is removed.
func Clean(appNames ...string) {
	if len(appNames) == 0 { // Remove whole workdir by default
		if Verbose {
			fmt.Println("Removing workdir at", workdir, "...")
		}

		if !DryRun { // Not if dry-running
			os.RemoveAll(workdir) // Remove /tmp/eac
		}

		if Verbose {
			fmt.Println("Removal finished.")
		}
	} else { // Remove specific apps only
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
	}
}
