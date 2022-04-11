package eac

import (
	"fmt"

	"github.com/thetillhoff/eac/pkg/eac/internal/apps"
	"github.com/thetillhoff/eac/pkg/eac/internal/dotProfile"
)

func ConfigureApp(appName string, dryRun bool) error {
	var (
		err          error
		app          = apps.Apps[appName]
		pathContains = false
	)

	if app.PathAddition != "" { // TODO Add to path if configured - only if not exists

		if Verbose {
			fmt.Println("Checking if envPath needs to be extended...")
		}

		pathContains, err = dotProfile.PathContains(app.PathAddition)
		if err != nil {
			return err
		}

		if Verbose {
			if pathContains {
				fmt.Println("Check finished.")
			}
		}

		if pathContains {
			if Verbose {
				fmt.Println("Skipped adding to envPath.")
			}
		} else {
			if Verbose {
				fmt.Println("Adding to envPath...")
			}
			if !dryRun {
				err = dotProfile.AddToPath(app.PathAddition)
				if err != nil {
					return err
				}
			}
			if Verbose {
				fmt.Println("EnvPath addition finished.")
			}
		}
	}

	if len(app.ProfileVariables) > 0 {
		for _, profileVariable := range app.ProfileVariables {
			if Verbose {
				fmt.Println("Checking if profile variable needs to be added...")
			}

			pathContains, err = dotProfile.ProfileContainsEnvVariable(profileVariable.Key, profileVariable.Value)
			if err != nil {
				return err
			}

			if Verbose {
				if pathContains {
					fmt.Println("Check finished.")
				}
			}

			if pathContains {
				if Verbose {
					fmt.Println("Skipped adding profile variable.")
				}
			} else {
				if Verbose {
					fmt.Println("Adding profile variable...")
				}
				if !dryRun {
					err = dotProfile.AddEnvVariableToProfile(profileVariable.Key, profileVariable.Value)
					if err != nil {
						return err
					}
				}
				if Verbose {
					fmt.Println("Addition of profile variable finished.")
				}
			}
		}
	}

	return err
}
