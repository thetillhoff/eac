package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thetillhoff/eac/pkg/apps"
	"github.com/thetillhoff/eac/pkg/logs"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the structure for a new app.",
	Long: `Create initial file-structure for new apps with specified names. Examples:
  eac create demo
	eac create demo1 demo2`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logs.ContinueOnError = conf.ContinueOnError
		logs.Verbose = conf.Verbose // needs to be done here, the other cmds pass it around
		flaggedPlatforms, err := cmd.Flags().GetStringSlice("platform")
		if err != nil {
			logs.Error("There was an error while reading the flag 'platform':", err)
		}

		flaggedGithubUser, err := cmd.Flags().GetString("githubUser")
		if err != nil {
			logs.Error("There was an error while reading the flag 'githubUser':", err)
		}

		createData := map[string]string{}
		if flaggedGithubUser != "" {
			createData["githubUser"] = flaggedGithubUser
		}

		apps.Create(args, flaggedPlatforms, conf.AppsDirPath, conf.Verbose, createData)
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

	createCmd.Flags().StringSliceP("platform", "p", []string{}, "Only create demo files for specified platforms. Valid options are ["+strings.Join(apps.ValidPlatforms(), "|")+"]")

	createCmd.Flags().String("githubUser", "", "Create default files specifically for github.")

	//TODO --fail-if-app-exists: bool
	//TODO --fail-if-platform-exists: bool

	viper.BindPFlags(createCmd.Flags())
	viper.UnmarshalExact(&conf)
}
