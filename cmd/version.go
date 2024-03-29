/*
Copyright © 2023 Till Hoffmann <till@thetillhoff.de>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev" // This is just the default. The actual value is injected at compiletime

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of eac",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
