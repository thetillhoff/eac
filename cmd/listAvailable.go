package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	eac "github.com/thetillhoff/eac/pkg/eac"
)

// listAvailableCmd represents the available command
var listAvailableCmd = &cobra.Command{
	Use:   "available",
	Short: "List all available apps.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// Passing flags & viper config
		eac.Verbose = verbose || viper.GetBool("verbose")

		appList, err := eac.ListAvailableApps()
		if err != nil {
			logger.Fatal(err.Error())
		}

		for _, appName := range appList {
			fmt.Println(appName)
		}

	},
}

func init() {
	listCmd.AddCommand(listAvailableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// availableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// availableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
