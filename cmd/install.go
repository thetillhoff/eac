package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/thetillhoff/eac/internal/app"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install and configure specified apps.",
	Long: `Install one or multiple apps, specified by a space-seperated list of names, f.e.
	eac install app1 app2 app3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		noConfigure, err := cmd.Flags().GetBool("no-configure")
		if err != nil {
			log.Fatalln(err)
		}

		for _, arg := range args {
			appItem := app.NewApp(arg)
			out, err := app.Install(appItem)
			out = strings.TrimSuffix(out, "\n")
			fmt.Println(out)
			if err != nil {
				log.Fatal(err)
			}

			if !noConfigure {
				out, err := app.Configure(appItem)
				out = strings.TrimSuffix(out, "\n")
				fmt.Println(out)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	installCmd.Flags().Bool("no-configure", false, "Don't run app configuration after their installation") //TODO implementation
	installCmd.Flags().BoolP("update", "u", false, "Update app versions before installation")              //TODO implementation
}
