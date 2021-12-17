package main

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	app     _Application
)

var rootCmd = &cobra.Command{
	Use: "pismo",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		app, err = initApplication(cmd)
		return err
	},
}

func main() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pismo.toml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose mode")
	cobra.CheckErr(rootCmd.Execute())
}
