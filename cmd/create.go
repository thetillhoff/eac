package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

var demoFiles = map[string]map[string]string{
	"linux": {
		"configure.sh": `#/bin/sh
echo "This script is called to configure of app %v"`,
		"getLatestVersion.sh": `#/bin/sh
echo "This script is called to getLatestVersion of app %v"`,
		"getLocalVersion.sh": `#/bin/sh
echo "This script is called to getLocalVersion of app %v"`,
		"install.sh": `#/bin/sh
echo "This script is called to install of app %v"`,
		"uninstall.sh": `#/bin/sh
echo uninstall %v`,
	},
	"darwin": {
		"configure.sh": `#/bin/sh
echo "This script is called to configure of app %v"`,
		"getLatestVersion.sh": `#/bin/sh
echo "This script is called to getLatestVersion of app %v"`,
		"getLocalVersion.sh": `#/bin/sh
echo "This script is called to getLocalVersion of app %v"`,
		"install.sh": `#/bin/sh
echo "This script is called to install of app %v"`,
		"uninstall.sh": `#/bin/sh
echo "This script is called to uninstall of app %v"`,
	},
	"windows": {
		"configure.ps1":        `Write-Host "This script is called to configure of app %v"`,
		"getLatestVersion.ps1": `Write-Host "This script is called to getLatestVersion of app %v"`,
		"getLocalVersion.ps1":  `Write-Host "This script is called to getLocalVersion of app %v"`,
		"install.ps1":          `Write-Host "This script is called to install of app %v"`,
		"uninstall.ps1":        `Write-Host "This script is called to uninstall of app %v"`,
	},
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the structure for a new app.",
	Long: `Create initial file-structure for a new app with specified name. Example:
  eac create demo`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appNames := args
		flagged_platforms, err := cmd.Flags().GetStringSlice("platform")
		if err != nil {
			log.Fatalln(err)
		}

		platforms := []string{}

		if len(flagged_platforms) == 0 { // for all platforms
			for platform := range demoFiles {
				platforms = append(platforms, platform)
			}
		} else { // for specific platforms
			for _, platform := range flagged_platforms { // validate specified platforms
				if !isValidPlatform(platform) {
					log.Fatalln("Unkown platform: " + platform)
				}
				platforms = append(platforms, platform)
			}
		}

		if _, err := os.Stat(appsPath); os.IsNotExist(err) {
			err := os.Mkdir(appsPath, os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Created '" + appsPath + "' folder.")
		} //TODO should this fail and point to 'init' command?

		for _, appName := range appNames {

			appPath := path.Join(appsPath, appName)
			if _, err := os.Stat(appPath); os.IsNotExist(err) { // if folder doesn't exist yet
				err := os.Mkdir(appPath, os.ModePerm)
				if err != nil {
					log.Fatalln(err)
				}
				log.Println("Created '" + appPath + "' folder.")
			} else if err == nil { // if folder does exist
				fmt.Println("App '" + appName + "' does already exist.")
			} else {
				log.Fatalln(err)
			}

			for _, platform := range platforms {
				platformPath := path.Join(appPath, platform)
				if _, err := os.Stat(platformPath); os.IsNotExist(err) {
					err := os.Mkdir(platformPath, os.ModePerm) // ignore errors
					if err != nil {
						log.Fatalln(err)
					}
					log.Println("Created '" + platformPath + "' folder.")
				} else if err == nil {
					fmt.Println("Platform '" + platform + "' for app '" + appName + "' does already exist.")
				} else {
					log.Fatalln(err)
				}

				platformDemoFiles := demoFiles[platform]

				for filename, fileContent := range platformDemoFiles {
					f, err := os.Create(path.Join(platformPath, filename))
					if err != nil {
						log.Fatalln(err)
					}
					defer f.Close()

					w := bufio.NewWriter(f)

					_, err = fmt.Fprintf(w, fileContent, appName)
					if err != nil {
						log.Fatalln(err)
					}

					w.Flush()
					log.Println("Created '" + path.Join(platformPath, filename) + "' file.")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	validPlatforms := make([]string, 0, len(demoFiles))
	for platform := range demoFiles {
		validPlatforms = append(validPlatforms, platform)
	}

	createCmd.Flags().StringSliceP("platform", "p", []string{}, "Only create demo files for specified platforms. Valid options are ["+strings.Join(validPlatforms, "|")+"]")

	//TODO --fail-if-app-exists: bool
	//TODO --fail-if-platform-exists: bool
}

func isValidPlatform(potentialPlatform string) bool {
	for platform := range demoFiles {
		if platform == potentialPlatform {
			return true
		}
	}
	return false
}
