package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eac/pkg/apps"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available apps.",
	Long: `List all available apps, with their respective version. Call with
	eac list`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		noVersion, err := cmd.Flags().GetBool("no-version")
		if err != nil {
			log.Fatalln("There was an error while reading the flag 'no-version':\n" + err.Error())
		}
		seperator, err := cmd.Flags().GetString("seperator")
		if err != nil {
			log.Fatalln("There was an error while reading the flag 'seperator':\n" + err.Error())
		}
		apps.List(appsDirPath, versionsFilePath, noVersion, seperator)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listCmd.Flags().Bool("no-version", false, "Don't show versions of apps.")
	listCmd.Flags().String("seperator", "\n", "Change seperator, default to \\n.")
}
