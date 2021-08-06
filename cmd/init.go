package cmd

import (
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eac/pkg/apps"
	"github.com/thetillhoff/eac/pkg/logs"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new apps directory with eac as demo example.",
	Long: `Create initial file-structure for new apps and add self (=='eac') as first entry. Example:
  eac init`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logs.ContinueOnError = continueOnError
		flaggedPlatforms, err := cmd.Flags().GetStringSlice("platform")
		if err != nil {
			logs.Err("There was an error while reading the flag 'platform':", err)
		}

		// creating versionsFile when not exists
		if _, err := os.Stat(versionsFilePath); os.IsNotExist(err) { // check if file exists; if not exists
			versionsFile, err := os.Create(versionsFilePath)
			if err != nil {
				logs.Err("Couldn't create versionsFile at '"+versionsFilePath+"':", err)
			} else {
				logs.Success("Created versionsFile.")
			}
			versionsFile.Close()
		} else { // when file does exist
			logs.Info("versionsFile already exists.")
		}

		if _, err := os.Stat(appsDirPath); os.IsNotExist(err) {
			err := os.Mkdir(appsDirPath, os.ModePerm)
			if err != nil {
				logs.Err("Couldn't create appsDir at '"+appsDirPath+"':", err)
			}
			logs.Info("Created '" + appsDirPath + "' folder.")
		} else if err == nil {
			appsDir, err := os.Open(appsDirPath) // open appsDir to check if it's empty
			if err != nil {
				logs.Err("There was a problem opening appsDir at '"+appsDirPath+"':", err)
			}
			defer appsDir.Close()

			_, err = appsDir.Readdirnames(1)
			if err != io.EOF { // check if appsDir is empty
				if err == nil {
					logs.Warn("Folder '" + appsDirPath + "' isn't empty.")
				} else {
					logs.Warn("Folder '"+appsDirPath+"' isn't accessible;", err)
				}
			}
		} else {
			logs.Err("There was a problem while accessing appsDir at '"+appsDirPath+"':", err)
		}

		apps.Create([]string{"eac"}, flaggedPlatforms, shell, appsDirPath, verbose)
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

	initCmd.Flags().StringSliceP("platform", "p", []string{}, "Only create demo files for specified platforms. Valid options are ["+strings.Join(apps.ValidPlatforms(), "|")+"]")
}
