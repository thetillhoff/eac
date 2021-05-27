package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var eacAppFiles = map[string]map[string]string{
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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new apps directory with eac as demo example.",
	Long: `Create initial file-structure for new apps and add self (=='eac') as first entry. Example:
  eac init`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(appsPath); os.IsNotExist(err) {
			err := os.Mkdir(appsPath, os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Created '" + appsPath + "' folder.")
		} else if err == nil {
			log.Fatalln("Folder '" + appsPath + "' already exists.") //TODO don't fail here, but check if the folder is empty or not and fail or continue based on that.
		} else {
			log.Fatalln(err)
		}

		appName := "eac"

		appPath := path.Join(appsPath, appName)
		err := os.Mkdir(appPath, os.ModePerm) // since parent folder is either freshly created or checked for emptiness at this point, no further checking is required
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Created '" + appPath + "' folder.")

		for platform := range eacAppFiles {
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

				_, err = fmt.Fprint(w, fileContent)
				if err != nil {
					log.Fatalln(err)
				}

				w.Flush()
				log.Println("Created '" + path.Join(platformPath, filename) + "' file.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
