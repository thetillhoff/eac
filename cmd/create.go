package cmd

import (
	"strings"

	"github.com/spf13/cobra"
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
		logs.ContinueOnError = continueOnError
		flaggedPlatforms, err := cmd.Flags().GetStringSlice("platform")
		if err != nil {
			logs.Err("There was an error while reading the flag 'platform':", err)
		}

		apps.Create(args, flaggedPlatforms, appsDirPath, verbose)
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

	//TODO --fail-if-app-exists: bool
	//TODO --fail-if-platform-exists: bool
}
