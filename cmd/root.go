package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{Use: "qnapcc"}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default . and $HOME/.qnapcc.yml)")
	//	rootCmd.Flags().StringP("url", "u", "http://192.168.1.1:8080", "QNAP endpoint URL")

	//	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".qnapcc")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("./")
	}
	viper.SetConfigType("yml")

	// Envirnomental Values
	// viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Run() error {
	return rootCmd.Execute()
}
