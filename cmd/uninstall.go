package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thetillhoff/eac/pkg/apps"
	"github.com/thetillhoff/eac/pkg/logs"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall specified apps.",
	Long: `Uninstall one or multiple apps, specified by a space-seperated list of names, f.e.
	eac update app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logs.ContinueOnError = conf.ContinueOnError
		logs.Verbose = conf.Verbose // needs to be done here, the other cmds pass it around

		apps.Uninstall(args, conf.AppsDirPath, conf.Verbose, conf.VersionsFilePath) // Uninstall apps
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.BindPFlags(uninstallCmd.Flags())
	viper.UnmarshalExact(&conf)

	//TODO: uninstallCmd.Flags().Bool("force", false, "Force uninstallation of *all* apps. Use with caution.")
}
