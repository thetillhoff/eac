package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specified app from eac tree.",
	Long: `Delete the file-structure for the specified app. Example call:
	eac delete demo`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appNames := args
		flagged_platforms, err := cmd.Flags().GetStringSlice("platform")
		if err != nil {
			log.Fatalln(err)
		}

		for _, appName := range appNames { // for all apps
			appPath := path.Join(appsPath, appName)
			if _, err := os.Stat(appPath); os.IsNotExist(err) { // if folder doesn't exist
				log.Fatalln("Folder '" + appPath + "' for app '" + appName + "' doesn't exist.")
			} else if err != nil {
				log.Fatalln(err)
			}

			if len(flagged_platforms) == 0 { // delete whole app
				err := os.RemoveAll(appPath)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println("Deleted '" + appPath + "' folder.")
			} else { // for specific platforms
				platforms := []string{}

				for _, platform := range flagged_platforms { // validate specified platforms
					if !isValidPlatform(platform) { // references method 'isValidPlatform' in 'create.go'
						log.Fatalln("Unkown platform: " + platform)
					}
					platforms = append(platforms, platform)

					for _, platform := range platforms {
						platformPath := path.Join(appPath, platform)
						if _, err := os.Stat(platformPath); os.IsNotExist(err) { // if folder doesn't exist
							log.Fatalln("Folder '" + platformPath + "' for platform '" + platform + "' for app '" + appName + "' doesn't exist.")
						} else if err != nil {
							log.Fatalln(err)
						}
						err := os.RemoveAll(platformPath)
						if err != nil {
							log.Fatalln(err)
						}
						fmt.Println("Deleted '" + platformPath + "' folder.")
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	validPlatforms := make([]string, 0, len(demoFiles)) // references variable 'demoFiles' in create.go
	for platform := range demoFiles {
		validPlatforms = append(validPlatforms, platform)
	}

	deleteCmd.Flags().StringSliceP("platform", "p", []string{}, "Only delete files for specified platforms. Valid options are ["+strings.Join(validPlatforms, "|")+"]")
}
