package cmd

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:           "anitya",
	Short:         "A anitya CLI written in Go.",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.anitya.yaml)")
	RootCmd.PersistentFlags().StringP("url", "", "https://anitya.example.com/api", "anitya endpoint URL")

	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".anitya")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}

func newDefaultClient() (*Client, error) {
	endpointURL := viper.GetString("url")
	httpClient := &http.Client{}
	userAgent := fmt.Sprintf("anitya/%s (%s)", Version, runtime.Version())
	return newClient(endpointURL, httpClient, userAgent)
}
