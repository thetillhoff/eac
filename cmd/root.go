package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/spf13/viper"
)

var (
	err     error
	cfgFile string
	logger  *zap.Logger
	verbose bool
	dryRun  bool // Only used byg installCmd and cleanCmd
	latest  bool // Only used by installCmd
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eac",
	Short: "eac installs, configures, upgrades, downgrades and uninstalls applications",
	Long: `eac is a package manager for applications that are not available in common alternatives
and for applications whose users need to switch between their versions often.`,
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eac)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "increase verbosity")

	// Configuring logger

	loggerConfig := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding:         "console", // one of ["json","console"]
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseColorLevelEncoder,
		},
	}

	logger, err = loggerConfig.Build()
	if err != nil {
		log.Fatalln(err)
	}
	// logger.Info("test-info")
	// logger.Warn("test-warning")
	// logger.Error("test-error")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var (
		home       string
		actualUser *user.User
	)
	if cfgFile != "" { // If cfgFile is set via flag (`--cfgFile <path>`)
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Since `eac` requires sudo most of the time, we need to work around that to get the actual user home
		sudoUser := os.Getenv("SUDO_USER")
		if sudoUser == "" { // If no sudo is involved
			home, err = os.UserHomeDir() // Get user's $HOME.
			if err != nil {
				log.Fatalln(err)
			}
		} else { // `eac` was ran with sudo
			actualUser, err = user.Lookup(sudoUser)
			if err != nil {
				log.Fatalln(err)
			}
			home = actualUser.HomeDir
		}

		// Search config in home directory with name ".eac" (without extension).
		viper.AddConfigPath(home)   // Set containing folder
		viper.SetConfigType("yaml") // Set config file format
		viper.SetConfigName(".eac") // Set config file name
	}

	viper.AutomaticEnv() // Read in environment variables that match the available config options

	// If a config file is found, read it in.
	err = viper.ReadInConfig()
	if err == nil { // If config file was found
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
