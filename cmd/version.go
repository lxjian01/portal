package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	globalConf "portal/global/config"
)

func init() {
	rootCmd.AddCommand()
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		conf := globalConf.GetAppConfig()
		fmt.Printf("version is %s \n", conf.Version)
	},
}