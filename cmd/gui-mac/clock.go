package main

import (
	"context"
	"fmt"
	"os"

	"github.com/caseymrm/menuet"
	"github.com/mniak/qisno/internal/utils"
)

func (a _Application) generateClockMenuItems(ctx context.Context) []menuet.MenuItem {
	info, err := a.ClockManager.Query(ctx)
	fmt.Println(info.FirstStartTime)
	if err != nil {
		fmt.Fprintln(os.Stderr, "falha na consulta", err)
		return []menuet.MenuItem{
			{
				Text: "Ponto: falha na consulta",
			},
			{
				Text:     "Marcar Ponto",
				FontSize: 14,
				Clicked:  a.clock,
			},
		}
	}
	if info.Empty {
		return []menuet.MenuItem{
			{
				Text: "Ponto: não marcado",
			},
			{
				Text:     "Marcar Entrada",
				FontSize: 14,
				Clicked:  a.clock,
			},
		}
	}

	if info.Running {
		return []menuet.MenuItem{
			{
				Text: fmt.Sprintf("Ponto: %s-??:?? (%s)",
					info.FirstStartTime.Format("15:04"),
					utils.FormatDuration(info.TotalTimeToday)),
			},
			menuet.MenuItem{
				Text:     "Marcar Saída",
				FontSize: 14,
				Clicked:  a.clock,
			},
		}
	}

	return []menuet.MenuItem{
		{
			Text: fmt.Sprintf("Ponto: %s-%s (%s)",
				info.FirstStartTime.Format("15:04"),
				info.LastEndTime.Format("15:04"),
				utils.FormatDuration(info.TotalTimeToday)),
		},
		{
			Text:     "Reiniciar Ponto",
			FontSize: 14,
			Clicked:  a.clock,
		},
	}
}

func (a _Application) clock() {
	ctx := newContext()
	err := a.ClockManager.Clock(ctx)
	if err != nil {
		a.ShowMessage("Falha ao bater o ponto!")
	}
}
