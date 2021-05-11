package cmd

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sopre",
	Short: "SOPRE CLI to operate GUI",
	Long:  `With SOPRE CLI you can start / install your local sopre model`,
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
	//if err := viper.ReadInConfig(); err == nil {
	//}

	err := viper.ReadInConfig()
	if err != nil {
		home, err := homedir.Dir()

		cobra.CheckErr(err)
		errmsg := fmt.Sprintf("ERROR: Config not found. Created Config at %s/.sopre.yml", home)
		color.Red(errmsg)
		createConfig()
		initConfig()
	}

}

func createConfig() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
sopre-qas-path: "\\\\i90256\\Quintiq 609\\"
sopre-qas-binary: "QAS-Setup-64-6_0_9.exe"
sopre-install-path: "C:\\Program Files\\Quintiq\\Quintiq 60\\bin"
sopre-properties-path: "C:\\Program Files\\Quintiq\\Quintiq 60\\etc"
sopre-gui-path:
  d:
    ap:
      designer: "DEV_AP_60_QClient"
      client: "DEV_AP_60_QThinClient"
      editor: "DEV_AP_60_QThinClientEditor"
    ep:
      designer: "DEV_EP_60_QClient"
      client: "DEV_EP_60_QThinClient"
      editor: "DEV_EP_60_QThinClientEditor"
  bc:
    ep:
      designer: "BCTEST_EP_60_QClient"
      client: "BCTEST_EP_60_QThinClient"
      editor: "BCTEST_EP_60_QThinClientEditor"
    global:
      designer: "BCTEST_GLOBAL_60_QClient"
      client: "BCTEST_GLOBAL_60_QThinClient"
      editor: "BCTEST_GLOBAL_60_QThinClientEditor"
    le:
      designer: "BCTEST_LE_60_QClient"
      client: "BCTEST_LE_60_QThinClient"
      editor: "BCTEST_LE_60_QThinClientEditor"
    plst:
      designer: "BCTEST_PLST_60_QClient"
      client: "BCTEST_PLST_60_QThinClient"
      editor: "BCTEST_PLST_60_QThinClientEditor"
`)

	_ = viper.ReadConfig(bytes.NewBuffer(yamlExample))

	home, err := homedir.Dir()

	cobra.CheckErr(err)
	// Search config in home directory with name ".sopre" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".sopre")
	viper.SetConfigType("yaml")

	viper.WriteConfigAs(path.Join(home, ".sopre.yml"))

}
