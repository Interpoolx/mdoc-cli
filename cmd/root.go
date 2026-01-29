package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	verbose   bool
	Version   string
	Commit    string
	BuildTime string
)

var rootCmd = &cobra.Command{
	Use:   "mdoc-cli",
	Short: "A high-performance markdown CLI tool",
	Long: `mdoc-cli provides fast and efficient markdown processing, 
conversion, and validation capabilities.

Built with Go for exceptional performance and cross-platform support.
Visit https://github.com/interpoolx/mdoc-cli for more information.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default searches: ./mdoc-cli.yaml, ~/.mdoc-cli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,
		"verbose output")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("$HOME/.config/mdoc-cli")
		viper.AddConfigPath("/etc/mdoc-cli")
		viper.SetConfigName("mdoc-cli")
	}

	viper.SetEnvPrefix("MDOC_CLI")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil && verbose {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
