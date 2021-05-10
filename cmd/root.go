package cmd

import (
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"runtime"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sopre",
	Short: "SOPRE CLI to operate GUI",
	Long:  `With SOPRE CLI you can start / install your local sopre model`,
	PreRun: func(cmd *cobra.Command, args []string) {

		runos := runtime.GOOS

		if runos == "windows" {
			color.Red("ERROR: This software only runs on Window Machines\n\n")
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sopre.yaml)")
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
		// Search config in home directory with name ".sopre" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".sopre")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
	}
}
