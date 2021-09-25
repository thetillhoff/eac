package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eac/pkg/logs"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	continueOnError  bool
	verbose          bool
	appsDirPath      string //TODO add eac to $PATH (don't forget to edit helptext for it)
	versionsFilePath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "eac",
	// Version: "0.0.0", // TODO set proper version -> could be done with github action prior to build.
	Short: "eac installs, configures, upgrads, downgrades and uninstalls applications.",
	Long:  `eac installs, configures, upgrads, downgrades and uninstalls applications.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file location, defaults to $HOME/.eac.yaml")
	rootCmd.PersistentFlags().StringVarP(&appsDirPath, "appsDirPath", "a", appsDirPath, "Override location of apps, defaults to ./apps/")
	rootCmd.PersistentFlags().BoolVar(&continueOnError, "continue-on-error", false, "Continue with other tasks even on failures, defaults to false")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Display more information during command execution, defaults to false")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Set the default paths
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		logs.Err("Can't retrieve userHomeDir", err)
	}
	eacDir := path.Join(userHomeDir, ".eac")
	appsDirPath = path.Join(eacDir, "apps")
	versionsFilePath = path.Join(eacDir, "versions.yaml")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".eac" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".eac")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logs.Info("Using config file:", viper.ConfigFileUsed())
	}
}
