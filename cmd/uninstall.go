package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thetillhoff/eac/pkg/apps"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall specified apps.",
	Long: `Uninstall one or multiple apps, specified by a space-seperated list of names, f.e.
	eac update app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apps.Uninstall(args, shell, appsDirPath, continueOnError, verbose) // Uninstall apps
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

	//TODO: uninstallCmd.Flags().Bool("force", false, "Force uninstallation of *all* apps. Use with caution.")
}
