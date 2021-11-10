package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thetillhoff/eac/pkg/apps"
	"github.com/thetillhoff/eac/pkg/logs"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install and configure specified apps.",
	Long: `Install one or multiple apps, specified by a space-seperated list of names, f.e.
	eac install app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logs.ContinueOnError = conf.ContinueOnError

		apps.Install(args, conf) // Install apps
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	installCmd.Flags().Bool("no-config", false, "Don't run app configuration after their installation")
	installCmd.Flags().BoolP("update", "u", false, "Update app versions before installation")

	installCmd.Flags().BoolP("latest", "l", false, "Install latest versions of apps, no matter which versions are specified anywhere.")

	viper.BindPFlags(installCmd.Flags())
	viper.UnmarshalExact(&conf)

	//TODO add flag for multiple versionsFilePaths that override each other
}
