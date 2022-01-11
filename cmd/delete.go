package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thetillhoff/eac/pkg/apps"
	"github.com/thetillhoff/eac/pkg/logs"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specified app from eac tree.",
	Long: `Delete the file-structure for the specified apps. Examples:
	eac delete demo
	eac delete demo1 demo2`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logs.ContinueOnError = conf.ContinueOnError
		flaggedPlatforms, err := cmd.Flags().GetStringSlice("platform")
		if err != nil {
			logs.Error("There was an error while reading the flag 'platform':", err)
		}

		apps.Delete(args, flaggedPlatforms, conf.AppsDirPath, conf.Verbose)
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

	deleteCmd.Flags().StringSliceP("platform", "p", []string{}, "Only delete files for specified platforms. Valid options are ["+strings.Join(apps.ValidPlatforms(), "|")+"]")

	viper.BindPFlags(deleteCmd.Flags())
	viper.UnmarshalExact(&conf)
}
