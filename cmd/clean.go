package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	eac "github.com/thetillhoff/eac/pkg/eac"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove all cached files for specified apps (Default: all apps)",
	Long: `Remove all cached files for the specified apps, i.e. all downloaded files. If no app is specified, all cached files are removed. Does not affect installed apps. Example:
	eac clean
	eac clean kubectl
	eac clean kubectl terraform`,
	Run: func(cmd *cobra.Command, args []string) {

		// Passing flags & viper config
		eac.Verbose = verbose || viper.GetBool("verbose")
		eac.DryRun = dryRun || viper.GetBool("dry-run")

		eac.Clean(args...)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cleanCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "prepare installation, but skip installation")
}
