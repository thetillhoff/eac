package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eac/pkg/apps"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for newer version of specified apps.",
	Long: `Check for newer version of one or multiple apps, specified by a space-seperated list of names, f.e.
	eac update app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			log.Fatalln("There was an error while reading the flag 'dry-run':\n" + err.Error())
		}
		skipLocal, err := cmd.Flags().GetBool("skip-local")
		if err != nil {
			log.Fatalln("There was an error while reading the flag 'skip-local':\n" + err.Error())
		}
		apps.Update(args, shell, appsDirPath, continueOnError, versionsFilePath, dryRun, skipLocal)
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

	updateCmd.Flags().Bool("dry-run", false, "Output potential upgrades, but don't make them.")
	updateCmd.Flags().Bool("skip-local", false, "Skip checking local version.")
	//TODO --dry-run: don't store/update version in settings.yaml. Only makes sense when combined with some install option. Maybe `install --update --no-save-version` or `install --latest --no-save-version`?
}
