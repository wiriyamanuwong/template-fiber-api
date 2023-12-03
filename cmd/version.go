/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	AppName   string = "example"
	Version   string = "0"
	TS        string
	Timestamp time.Time
	Build     string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "api version",
	Long:  `show api version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("App Name:", AppName)
		fmt.Println("App Version:", Version)
		// viper.SetEnvPrefix("FB")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	if TS != "" {
		Timestamp, _ = time.Parse(time.RFC3339, TS)
	} else {
		Timestamp = time.Now()
	}
	if Build == "" {
		Build = "dev"
	}

	Version = fmt.Sprintf("%.4d-%.2d-%.2d.%s", Timestamp.Year(), Timestamp.Month(), Timestamp.Day(), Build)
	// log.Info().Msg("Version: " + Version)

	viper.SetDefault("appname", AppName)
	viper.SetDefault("version", Version)
	viper.SetDefault("Timestamp", Timestamp)

}
