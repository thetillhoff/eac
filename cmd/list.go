package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thetillhoff/eac/pkg/apps"
	"github.com/thetillhoff/eac/pkg/logs"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available apps.",
	Long: `List all available apps, with their respective version. Call with
	eac list`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logs.ContinueOnError = conf.ContinueOnError
		logs.Verbose = conf.Verbose // needs to be done here, the other cmds pass it around

		apps.PrintList(conf)
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

	viper.BindPFlags(listCmd.Flags())
	viper.UnmarshalExact(&conf)
}
