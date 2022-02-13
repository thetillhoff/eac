package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thetillhoff/eac/pkg/apps"
	"github.com/thetillhoff/eac/pkg/logs"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for newer version of specified apps.",
	Long: `Check for newer version of one or multiple apps, specified by a space-seperated list of names, f.e.
	eac update app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logs.ContinueOnError = conf.ContinueOnError
		logs.Verbose = conf.Verbose // needs to be done here, the other cmds pass it around
		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			logs.Error("There was an error while reading the flag 'dry-run':", err)
		}
		conf.DryRun = dryRun
		skipLocal, err := cmd.Flags().GetBool("skip-local")
		if err != nil {
			logs.Error("There was an error while reading the flag 'skip-local':", err)
		}
		conf.SkipLocal = skipLocal

		apps.Update(args, conf)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	updateCmd.PersistentFlags().BoolVar(&conf.UpdateConfig.DryRun, "dry-run", conf.UpdateConfig.DryRun, "Output potential upgrades, but don't make them.")
	updateCmd.PersistentFlags().BoolVar(&conf.UpdateConfig.SkipLocal, "skip-local", conf.UpdateConfig.SkipLocal, "Skip checking local version.")

	viper.BindPFlags(updateCmd.Flags())
	viper.UnmarshalExact(&conf)

	//TODO --dry-run: don't store/update version in settings.yaml. Only makes sense when combined with some install option. Maybe `install --update --no-save-version` or `install --latest --no-save-version`?
}
