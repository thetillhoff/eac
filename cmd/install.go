package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	eac "github.com/thetillhoff/eac/pkg/eac"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install and configure specified apps",
	Long: `Install one or multiple apps, specified by a space-separated list of names with optional versions. Examples:
	eac install
  eac install kubectl@1.2.3
	eac install kubectl terraform`,
	Run: func(cmd *cobra.Command, args []string) {

		// Passing flags & viper config
		eac.Verbose = verbose || viper.GetBool("verbose")
		eac.DryRun = dryRun || viper.GetBool("dry-run")

		if len(args) != 0 { // If apps/versions are already specified explicitly
			viper.Set("apps", args)
		} // Else no explicit apps/versions are set use the config file instead

		if len(viper.GetStringSlice("apps")) == 0 { // Neither args nor `~/.eac` specify any apps
			logger.Fatal("No apps specified")
		}

		for _, appName := range viper.GetStringSlice("apps") {
			version := ""
			if strings.Contains(appName, "@") { // If appName contains '@'
				parts := strings.Split(appName, "@")
				if len(parts) > 2 {
					logger.Fatal("Invalid app " + appName)
				}
				appName = parts[0]
				if !latest && !viper.GetBool("latest") { // If latest is not set on global level
					version = parts[1]
				}
			}
			if version == "" {
				version, err = eac.GetLatestVersion(appName)
				if err != nil {
					logger.Fatal(err.Error())
				}
			}

			err = eac.InstallApp(appName, version)
			if err != nil {
				logger.Fatal(err.Error())
			}

			// err = eac.ConfigureApp(appName, dryRun)
			// if err != nil {
			// 	logger.Fatal(err.Error())
			// }
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

	installCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "prepare installation, but skip installation")
	installCmd.PersistentFlags().BoolVar(&latest, "latest", false, "install latest version for all specified apps")
}
