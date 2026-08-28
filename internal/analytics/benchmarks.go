package analytics

import "traininganalysis/internal/model"

type Benchmark struct {
	Drill     string
	Target    float64
	Tolerance float64
}

func Compare(result model.DrillResult, b Benchmark) string {
	delta := result.Score - b.Target
	if delta >= -b.Tolerance {
		return "on-track"
	}
	if delta >= -2*b.Tolerance {
		return "watch"
	}
	return "needs-work"
}
func BenchmarkTable(results []model.DrillResult, benchmarks []Benchmark) map[string]string {
	out := map[string]string{}
	for _, r := range results {
		for _, b := range benchmarks {
			if r.Drill == b.Drill {
				out[r.AthleteID+":"+r.Drill] = Compare(r, b)
			}
		}
	}
	return out
}
func Improvement(before, after []model.DrillResult) float64 {
	a := ScoreAverage(before)
	b := ScoreAverage(after)
	return b - a
}
func Consistency(results []model.DrillResult) float64 {
	if len(results) < 2 {
		return 100
	}
	avg := ScoreAverage(results)
	sum := 0.0
	for _, r := range results {
		d := r.Score - avg
		sum += d * d
	}
	return 100 / (1 + sum/float64(len(results)))
}
