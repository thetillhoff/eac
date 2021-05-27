package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available apps.",
	Long: `List all available apps, with their respective version. Call with
	eac list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called") //TODO
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

	//TODO --no-version: don't list version
	//TODO --seperator: override newline char with something else (string, not char)
}
