package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eac/internal/app"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure specified apps.",
	Long: `Configure one or multiple apps, specified by a space-seperated list of names, f.e.
	eac configure app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			appItem := app.NewApp(arg)

			// test if get-local-version works (to check if app is already installed), if not, fail (with custom error?) //TODO: create custom error message
			_, err := app.GetLocalVersion(appItem)
			if err != nil {
				log.Fatal(err)
			}

			out, err := app.Configure(appItem)
			out = strings.TrimSuffix(out, "\n")
			fmt.Println(out)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
