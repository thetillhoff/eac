package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eac/internal/app"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for newer version of specified apps.",
	Long: `Check for newer version of one or multiple apps, specified by a space-seperated list of names, f.e.
	eac update app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			appItem := app.NewApp(arg)

			out, err := app.Update(appItem)
			out = strings.TrimSuffix(out, "\n")
			fmt.Println(out)
			if err != nil {
				log.Fatal(err)
			}
		}
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

	//TODO --no-save-version: don't store/update version in settings.yaml. Only makes sense when combined with some install option. Maybe `install --update --no-save-version` or `install --latest --no-save-version`?
	//TODO add --version for each app ...
}
