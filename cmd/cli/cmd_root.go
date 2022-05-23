package main

import (
	"github.com/mniak/qisno/internal/utils"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	app     _Application
)

var rootCmd = &cobra.Command{
	Use: "qisno",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		utils.SetExecDebug(flagVerbose)

		var err error
		app, err = initApplication(cmd)
		return err
	},
}

var flagVerbose bool

func main() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.qisno.toml)")
	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "Enable verbose mode")
	cobra.CheckErr(rootCmd.Execute())
}
