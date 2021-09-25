package cmd

import (
	"github.com/thetillhoff/eac/pkg/apps"
	"github.com/thetillhoff/eac/pkg/logs"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install and configure specified apps.",
	Long: `Install one or multiple apps, specified by a space-seperated list of names, f.e.
	eac install app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logs.ContinueOnError = continueOnError
		noConfigure, err := cmd.Flags().GetBool("no-configure")
		if err != nil {
			logs.Err("There was an error while reading the flag 'no-configure':", err)
		}
		update, err := cmd.Flags().GetBool("update")
		if err != nil {
			logs.Err("There was an error while reading the flag 'update':", err)
		}
		latest, err := cmd.Flags().GetBool("latest")
		if err != nil {
			logs.Err("There was an error while reading the flag 'latest':", err)
		}

		apps.Install(args, noConfigure, update, appsDirPath, versionsFilePath, latest, verbose) // Install apps
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

	installCmd.Flags().Bool("no-configure", false, "Don't run app configuration after their installation")
	installCmd.Flags().BoolP("update", "u", false, "Update app versions before installation")

	installCmd.Flags().BoolP("latest", "l", false, "Install latest versions of apps, no matter which versions are specified anywhere.")

	//TODO add flag for multiple versionsFilePaths
}
