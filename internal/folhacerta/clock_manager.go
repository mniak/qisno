package folhacerta

import (
	"context"
	"fmt"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/go-resty/resty/v2"
	"github.com/mniak/pismo/pkg/pismo"
)

type _FolhaCertaClockManager struct {
	config Config
}

func (c *_FolhaCertaClockManager) Query(ctx context.Context) (pismo.ClockInfo, error) {
	resp, err := resty.New().SetDebug(c.config.Verbose).R().
		SetContext(ctx).
		SetMultipartFormData(map[string]string{
			"tokenApp": c.config.Token,
		}).
		SetResult(CarregarDiaResponse{}).
		Post("https://app.folhacerta.com/App/CarregarDia")
	if err != nil {
		return pismo.ClockInfo{}, err
	}
	if !resp.IsSuccess() {
		return pismo.ClockInfo{}, fmt.Errorf("received invalid status code %d", resp.StatusCode())
	}
	result := resp.Result().(*CarregarDiaResponse)
	if !result.Success {
		return pismo.ClockInfo{}, fmt.Errorf("request failed: %s", result.Error)
	}

	var marcacoes []Marcacao
	linq.From(result.Dia.HorariosMarcacoes).
		Where(func(i interface{}) bool {
			x := i.(HorarioMarcacao)
			return x.Tipo == 1
		}).
		SelectMany(func(i interface{}) linq.Query {
			x := i.(HorarioMarcacao)
			return linq.From(x.Marcacoes).
				Where(func(i interface{}) bool {
					x := i.(Marcacao)
					return x.Tipo == 1 && x.ID != 44802715
				})
		}).
		OrderBy(func(i interface{}) interface{} {
			x := i.(Marcacao)
			d := parseDate(x.DataHora).UnixMilli()
			return d
		}).
		ToSlice(&marcacoes)
	info := pismo.ClockInfo{
		Empty:   len(result.Dia.HorariosMarcacoes) == 0,
		Running: len(marcacoes)%2 == 1,
	}
	if !info.Empty {
		info.FirstStartTime = parseDate(marcacoes[0].DataHora)
		// info.LastStartTime = parseDate(marcacoes[((len(marcacoes)-1)/2)*2].DataHora)
		info.LastEndTime = parseDate(marcacoes[len(marcacoes)-1].DataHora)
	}
	if !info.Empty {
		trabalhadas := parseDuration(result.Dia.Resumo.HorasTrabalhadas)
		if info.Running {
			info.TotalTimeToday = time.Now().Sub(info.FirstStartTime)
		} else {
			info.TotalTimeToday = trabalhadas
		}
	}
	return info, nil
}

func (c *_FolhaCertaClockManager) Clock(ctx context.Context) error {
	resp, err := resty.New().SetDebug(c.config.Verbose).R().
		SetContext(ctx).
		SetMultipartFormData(map[string]string{
			"tokenApp":    c.config.Token,
			"dispositivo": "Web",
			"tipo":        "5",
		}).
		SetResult(CarregarDiaResponse{}).
		Post("https://app.folhacerta.com/App/MarcarPonto")
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("received invalid status code %d", resp.StatusCode())
	}
	result := resp.Result().(*CarregarDiaResponse)
	if !result.Success {
		return fmt.Errorf("request failed: %s", result.Error)
	}
	return nil
}
