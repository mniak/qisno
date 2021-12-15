package folhacerta

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/go-resty/resty/v2"
	"github.com/mniak/pismo/domain"
)

type _FolhaCertaClockManager struct {
	config Config
}

func (c *_FolhaCertaClockManager) Query(ctx context.Context) (domain.ClockInfo, error) {
	resp, err := resty.New().SetDebug(true).R().
		SetContext(ctx).
		SetMultipartFormData(map[string]string{
			"tokenApp": c.config.Token,
		}).
		SetResult(CarregarDiaResponse{}).
		Post("https://app.folhacerta.com/App/CarregarDia")
	if err != nil {
		return domain.ClockInfo{}, err
	}
	if !resp.IsSuccess() {
		return domain.ClockInfo{}, fmt.Errorf("received invalid status code %d", resp.StatusCode())
	}
	result := resp.Result().(*CarregarDiaResponse)
	if !result.Success {
		return domain.ClockInfo{}, fmt.Errorf("request failed: %s", result.Error)
	}

	var marcacoes []Marcacao
	linq.From(result.Dia.HorariosMarcacoes).
		Where(func(i interface{}) bool {
			x := i.(HorarioMarcacao)
			return x.Tipo == 1
		}).
		SelectMany(func(i interface{}) linq.Query {
			x := i.(HorarioMarcacao)
			return linq.From(x.Marcacoes)
		}).
		ToSlice(&marcacoes)
	info := domain.ClockInfo{
		Empty:   len(result.Dia.HorariosMarcacoes) == 0,
		Running: len(marcacoes)%2 == 1,
	}
	if !info.Empty {
		info.FirstStartTime = parseDate(marcacoes[0].DataHora)
		info.LastEndTime = parseDate(marcacoes[len(marcacoes)-1].DataHora)
	}
	if !info.Empty {
		if info.Running {
			info.TotalTimeToday = time.Now().Sub(info.FirstStartTime)
		} else {
			info.TotalTimeToday = parseDuration(result.Dia.Resumo.HorasTrabalhadas)
		}
	}
	return info, nil
}

var dateRegex = regexp.MustCompile(`/Date\((\d+)\)/`)

func parseDate(s string) time.Time {
	subs := dateRegex.FindStringSubmatch(s)
	if len(subs) > 2 {
		fmt.Fprintf(os.Stderr, "failed to parse date %s\n", s)
		return time.Time{}
	}
	n, err := strconv.Atoi(subs[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse date %s\n", s)
		return time.Time{}
	}
	return time.Unix(int64(n)/1000, 0)
}

func parseDuration(s string) time.Duration {
	segments := strings.SplitN(s, ":", 2)
	if len(segments) < 2 {
		fmt.Fprintf(os.Stderr, "failed to parse duration %s\n", s)
		return 0
	}

	h, err := strconv.Atoi(segments[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse duration %s\n", s)
		return 0
	}

	m, err := strconv.Atoi(segments[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse duration %s\n", s)
		return 0
	}

	return time.Duration(h)*time.Hour + time.Duration(m)*time.Minute
}

func (c *_FolhaCertaClockManager) Clock(ctx context.Context) error {
	resp, err := resty.New().SetDebug(true).R().
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
