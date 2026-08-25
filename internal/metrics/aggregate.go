package metrics

import "traininganalysis/internal/model"

func Average(values []model.Metric) float64 {
	if len(values) == 0 {
		return 0
	}
	total := 0.0
	for _, v := range values {
		total += v.Value
	}
	return total / float64(len(values))
}
func ByUnit(values []model.Metric, unit string) []model.Metric {
	out := []model.Metric{}
	for _, v := range values {
		if v.Unit == unit {
			out = append(out, v)
		}
	}
	return out
}
