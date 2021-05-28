package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eac/pkg/apps"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the configuration of specified apps.",
	Long: `Validate the configuration of one or multiple apps, specified by a space-seperated list of names, f.e.
	eac validate app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flaggedPlatforms, err := cmd.Flags().GetStringSlice("platform")
		if err != nil {
			log.Fatalln("There was an error while reading the flag 'platform':\n" + err.Error())
		}

		apps.Validate(args, flaggedPlatforms, shell, appsDirPath, continueOnError)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	validateCmd.Flags().StringSliceP("platform", "p", []string{}, "Only create demo files for specified platforms. Valid options are ["+strings.Join(apps.ValidPlatforms(), "|")+"]")
}
