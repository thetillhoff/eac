package cmd

import (
	"os"
	"path"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eac/internal/config"
	"github.com/thetillhoff/eac/pkg/logs"

	"github.com/spf13/viper"
)

var (
	cfgFile       string
	VersionString string        = "0.0.0" // Override at compile time with "-ldflags '-X github.com/thetillhoff/eac/cmd.VersionString=1.2.3'"
	conf          config.Config = config.Config{
		ContinueOnError:  false,
		Platforms:        []string{runtime.GOOS},
		Verbose:          false,
		UserHomeDir:      "", // Set during init
		EacDirPath:       "", // Set during init
		AppsDirPath:      "", // Set during init
		VersionsFilePath: "", // Set during init
		ConfigureConfig:  config.ConfigureConfig{},
		CreateConfig: config.CreateConfig{
			GitHubUser: "",
		},
		DeleteConfig: config.DeleteConfig{},
		InitConfig:   config.InitConfig{},
		InstallConfig: config.InstallConfig{
			NoConfiguration: false,
			Update:          false,
			Latest:          false,
		},
		ListConfig: config.ListConfig{
			NoVersion: false,
			Seperator: "\n",
		},
		UninstallConfig: config.UninstallConfig{},
		UpdateConfig: config.UpdateConfig{
			DryRun:    false,
			SkipLocal: false,
		},
		ValidateConfig: config.ValidateConfig{},
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "eac",
	Version: VersionString,
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file location, defaults to $HOME/.eac.yaml")
	rootCmd.PersistentFlags().StringVarP(&conf.AppsDirPath, "apps-dir", "a", conf.AppsDirPath, "Override location of apps, defaults to ~/.eac/apps/")
	rootCmd.PersistentFlags().BoolVar(&conf.ContinueOnError, "continue-on-error", false, "Continue with other tasks even on failures, defaults to false")
	rootCmd.PersistentFlags().BoolVarP(&conf.Verbose, "verbose", "v", false, "Display more information during command execution, defaults to false")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var (
		err error
	)

	// Set the defaults
	conf.UserHomeDir, err = os.UserHomeDir()
	if err != nil {
		logs.Error("Can't retrieve userHomeDir", err)
	}
	conf.EacDirPath = path.Join(conf.UserHomeDir, ".eac")
	conf.AppsDirPath = path.Join(conf.UserHomeDir, ".eac", "apps")
	conf.VersionsFilePath = path.Join(conf.EacDirPath, "versions.yaml")

	// Search config in home directory with name ".eac.yaml".
	viper.AddConfigPath(conf.UserHomeDir)
	viper.SetConfigName(".eac.yaml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logs.Info("Using config file:", viper.ConfigFileUsed())
	}

	viper.BindPFlags(rootCmd.Flags())
	viper.UnmarshalExact(&conf)
}
