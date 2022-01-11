package cmd

import (
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thetillhoff/eac/internal/templates"
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
		logs.ContinueOnError = conf.ContinueOnError
		logs.Verbose = conf.Verbose // needs to be done here, the other cmds pass it around
		flaggedPlatforms, err := cmd.Flags().GetStringSlice("platform")
		if err != nil {
			logs.Error("There was an error while reading the flag 'platform':", err)
		}

		// Creating ~/.eac folder
		if _, err := os.Stat(conf.EacDirPath); os.IsNotExist(err) {
			err := os.Mkdir(conf.EacDirPath, os.ModePerm)
			if err != nil {
				logs.Error("Can't create folder at '"+conf.EacDirPath+"':", err)
			}
			logs.Info("Created folder at '" + conf.EacDirPath + "'.")
		} else if err == nil {
			dir, err := os.Open(conf.EacDirPath) // open folder to check if it's empty
			if err != nil {
				logs.Error("There was a problem opening folder at '"+conf.EacDirPath+"':", err)
			}
			defer dir.Close()

			_, err = dir.Readdirnames(1)
			if err != io.EOF { // check if appsDir is empty
				if err == nil {
					logs.Warn("Folder '" + conf.EacDirPath + "' isn't empty.")
				} else {
					logs.Warn("Folder '"+conf.EacDirPath+"' isn't accessible;", err)
				}
			}
		} else {
			logs.Error("There was a problem while accessing folder at '"+conf.EacDirPath+"':", err)
		}

		// creating versionsFile when not exists
		if _, err := os.Stat(conf.VersionsFilePath); os.IsNotExist(err) { // check if file exists; if not exists
			versionsFile, err := os.Create(conf.VersionsFilePath)
			if err != nil {
				logs.Error("Couldn't create versionsFile at '"+conf.VersionsFilePath+"':", err)
			} else {
				logs.Success("Created versionsFile.")
			}
			versionsFile.Close()
		} else { // when file does exist
			logs.Info("versionsFile already exists.")
		}

		// Creating appsDir
		if _, err := os.Stat(conf.AppsDirPath); os.IsNotExist(err) {
			err := os.Mkdir(conf.AppsDirPath, os.ModePerm)
			if err != nil {
				logs.Error("Couldn't create appsDir at '"+conf.AppsDirPath+"':", err)
			}
			logs.Info("Created '" + conf.AppsDirPath + "' folder.")
		} else if err == nil {
			appsDir, err := os.Open(conf.AppsDirPath) // open appsDir to check if it's empty
			if err != nil {
				logs.Error("There was a problem opening appsDir at '"+conf.AppsDirPath+"':", err)
			}
			defer appsDir.Close()

			_, err = appsDir.Readdirnames(1)
			if err != io.EOF { // check if appsDir is empty
				if err == nil {
					logs.Warn("Folder '" + conf.AppsDirPath + "' isn't empty.")
				} else {
					logs.Warn("Folder '"+conf.AppsDirPath+"' isn't accessible;", err)
				}
			}
		} else {
			logs.Error("There was a problem while accessing appsDir at '"+conf.AppsDirPath+"':", err)
		}

		// Creating folder for shared scripts
		if _, err := os.Stat(path.Join(conf.EacDirPath, "shared")); os.IsNotExist(err) {
			err := os.Mkdir(path.Join(conf.EacDirPath, "shared"), os.ModePerm)
			if err != nil {
				logs.Error("Couldn't create appsDir at '"+conf.EacDirPath+"':", err)
			}
			logs.Info("Created '" + path.Join(conf.EacDirPath, "shared") + "' folder.")
		} else if err == nil { // Folder is not empty
			logs.Info("Folder '" + path.Join(conf.EacDirPath, "shared") + "' is not empty.")
		} else if err == io.EOF { // Folder exists but is empty
			// Do nothing
		} else { // Other errors
			logs.Error("There was a problem while accessing folder at '"+path.Join(conf.EacDirPath, "shared")+"':", err)
		}

		// Creating shared scripts
		for _, sharedFile := range templates.GetSharedFiles() {
			if _, err := os.Stat(path.Join(conf.EacDirPath, "shared", strings.TrimPrefix(sharedFile, "shared-"))); os.IsNotExist(err) {
				templates.WriteTemplateToFile(path.Join(runtime.GOOS, sharedFile), new(interface{}), path.Join(conf.EacDirPath, "shared", strings.TrimPrefix(sharedFile, "shared-")))
				logs.Info("Created '" + path.Join(conf.EacDirPath, "shared", strings.TrimPrefix(sharedFile, "shared-")) + "' file.")
			} else {
				logs.Warn("Shared script at '" + path.Join(conf.EacDirPath, "shared", strings.TrimPrefix(sharedFile, "shared-")) + "' already exists.")
			}
		}

		// Creating eac app
		data := map[string]string{"githubUser": "thetillhoff"}
		apps.Create([]string{"eac"}, flaggedPlatforms, conf.AppsDirPath, conf.Verbose, data)
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

	viper.BindPFlags(initCmd.Flags())
	viper.UnmarshalExact(&conf)
}
