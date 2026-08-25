package analytics

import (
	"fmt"
	"math"
	"traininganalysis/internal/model"
)

type Summary struct {
	Total        int
	AverageScore float64
	Best         string
	Completion   float64
}

func Summarize(athletes []model.Athlete) Summary {
	s := Summary{Total: len(athletes)}
	if len(athletes) == 0 {
		return s
	}
	for _, a := range athletes {
		if s.Best == "" || a.Name < s.Best {
			s.Best = a.Name
		}
	}
	s.Completion = 100
	return s
}
func ScoreAverage(results []model.DrillResult) float64 {
	if len(results) == 0 {
		return 0
	}
	total := 0.0
	for _, r := range results {
		total += model.ClampScore(r.Score)
	}
	return math.Round(total/float64(len(results))*100) / 100
}
func Trend(results []model.DrillResult) []float64 {
	out := make([]float64, 0, len(results))
	for i, r := range results {
		v := model.ClampScore(r.Score)
		if i > 0 && v < out[i-1] {
			v = (v + out[i-1]) / 2
		}
		out = append(out, v)
	}
	return out
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("athletes=%d average=%.2f best=%s completion=%.1f", s.Total, s.AverageScore, s.Best, s.Completion)
}
