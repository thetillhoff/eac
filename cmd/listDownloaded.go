package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	eac "github.com/thetillhoff/eac/pkg/eac"
)

// listDownloadedCmd represents the downloaded command
var listDownloadedCmd = &cobra.Command{
	Use:   "downloaded",
	Short: "List all downloaded/cached app.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// Passing flags & viper config
		eac.Verbose = verbose || viper.GetBool("verbose")

		appList, err := eac.ListDownloadedApps(false)
		if err != nil {
			logger.Fatal(err.Error())
		}

		for _, entry := range appList {
			fmt.Println(entry)
		}

	},
}

func init() {
	listCmd.AddCommand(listDownloadedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
