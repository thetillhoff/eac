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
		flaggedPlatforms, err := cmd.Flags().GetStringSlice("platform")
		if err != nil {
			logs.Err("There was an error while reading the flag 'platform':", continueOnError, err)
		}

		if _, err := os.Stat(appsDirPath); os.IsNotExist(err) {
			err := os.Mkdir(appsDirPath, os.ModePerm)
			if err != nil {
				logs.Err("Couldn't create appsDir at '"+appsDirPath+"':", continueOnError, err)
			}
			logs.Info("Created '" + appsDirPath + "' folder.")
		} else if err == nil {
			appsDir, err := os.Open(appsDirPath) // open appsDir to check if it's empty
			if err != nil {
				logs.Err("There was a problem opening appsDir at '"+appsDirPath+"':", continueOnError, err)
			}
			defer appsDir.Close()

			_, err = appsDir.Readdirnames(1)
			if err != io.EOF { // check if appsDir is empty
				logs.Warn("Folder '"+appsDirPath+"' isn't empty or another problem occured while accessing it:", err)
			}
		} else {
			logs.Err("There was a problem while accessing appsDir at '"+appsDirPath+"':", continueOnError, err)
		}

		apps.Create([]string{"eac"}, flaggedPlatforms, shell, appsDirPath, true, verbose)
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
