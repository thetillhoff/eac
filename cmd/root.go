package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	shell            string
	continueOnError  bool
	appsDirPath      = path.Clean("apps") //TODO make this an absolute path, so eac can be added to $PATH (don't forget to edit helptext for it) // -> do this, when deciding where to put eac on linux systems, windows systems and their corresponding apps-dir
	versionsFilePath = path.Clean("versions.yaml")
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "eac",
	Version: "0.0.0",
	Short:   "eac installs, configures, upgrads, downgrades and uninstalls applications.",
	Long:    `eac installs, configures, upgrads, downgrades and uninstalls applications.`,
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file locatio, defaults to $HOME/.eac.yaml")
	rootCmd.PersistentFlags().StringVarP(&shell, "shell", "s", "", "Override shell for all apps, defaults to /bin/sh")
	rootCmd.PersistentFlags().StringVarP(&appsDirPath, "appsDirPath", "a", appsDirPath, "Override location of apps, defaults to ./apps/") //TODO: implementation
	rootCmd.PersistentFlags().BoolVar(&continueOnError, "continue-on-error", false, "Continue with other tasks even on failures, defaults to false")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

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
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
