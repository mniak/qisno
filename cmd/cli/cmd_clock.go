package main

import (
	"fmt"

	"github.com/mniak/qisno/internal/utils"
	"github.com/spf13/cobra"
)

var cmdClock = &cobra.Command{
	Use: "clock",
	Run: func(cmd *cobra.Command, args []string) {
		info, err := app.ClockManager.Query(newContext())
		handle(err, "falha na consulta")

		if info.Empty {
			fmt.Println("Ponto não marcado")
			return
		}

		if info.Running {
			fmt.Println("Ponto em aberto")
			fmt.Printf("Iniciado em %s\n", info.FirstStartTime.Format("15:04"))
			fmt.Printf("Total: %s\n", utils.FormatDuration(info.TotalTimeToday))
			return
		}

		fmt.Println("Ponto encerrado")
		fmt.Printf("Iniciado em %s\n", info.FirstStartTime.Format("15:04"))
		fmt.Printf("Encerrado em %s\n", info.LastEndTime.Format("15:04"))
		fmt.Printf("Total: %s\n", utils.FormatDuration(info.TotalTimeToday))
	},
}

func init() {
	rootCmd.AddCommand(cmdClock)
}
