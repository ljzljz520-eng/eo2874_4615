package analytics

import (
	"sort"
	"traininganalysis/internal/model"
)

func ScoreRange(results []model.DrillResult) (float64, float64) {
	if len(results) == 0 {
		return 0, 0
	}
	lo, hi := results[0].Score, results[0].Score
	for _, r := range results[1:] {
		if r.Score < lo {
			lo = r.Score
		}
		if r.Score > hi {
			hi = r.Score
		}
	}
	return lo, hi
}
func BestResult(results []model.DrillResult) (model.DrillResult, bool) {
	if len(results) == 0 {
		return model.DrillResult{}, false
	}
	best := results[0]
	for _, r := range results[1:] {
		if r.Score > best.Score {
			best = r
		}
	}
	return best, true
}
func WorstResult(results []model.DrillResult) (model.DrillResult, bool) {
	if len(results) == 0 {
		return model.DrillResult{}, false
	}
	worst := results[0]
	for _, r := range results[1:] {
		if r.Score < worst.Score {
			worst = r
		}
	}
	return worst, true
}
func CountAbove(results []model.DrillResult, threshold float64) int {
	n := 0
	for _, r := range results {
		if r.Score >= threshold {
			n++
		}
	}
	return n
}
func CountBelow(results []model.DrillResult, threshold float64) int {
	n := 0
	for _, r := range results {
		if r.Score < threshold {
			n++
		}
	}
	return n
}
func PercentAbove(results []model.DrillResult, threshold float64) float64 {
	if len(results) == 0 {
		return 0
	}
	return float64(CountAbove(results, threshold)) / float64(len(results)) * 100
}
func WeightedScore(results []model.DrillResult, weights map[string]float64) float64 {
	total, weight := 0.0, 0.0
	for _, r := range results {
		w := weights[r.Drill]
		if w <= 0 {
			w = 1
		}
		total += r.Score * w
		weight += w
	}
	if weight == 0 {
		return 0
	}
	return total / weight
}
func Reliability(results []model.DrillResult) float64 {
	if len(results) == 0 {
		return 0
	}
	valid := 0
	for _, r := range results {
		if r.Attempts > 0 && r.Score >= 0 {
			valid++
		}
	}
	return float64(valid) / float64(len(results)) * 100
}
func VolumeByAthlete(results []model.DrillResult) map[string]int {
	out := map[string]int{}
	for _, r := range results {
		out[r.AthleteID] += r.Attempts
	}
	return out
}
func AverageAttempts(results []model.DrillResult) float64 {
	if len(results) == 0 {
		return 0
	}
	n := 0
	for _, r := range results {
		n += r.Attempts
	}
	return float64(n) / float64(len(results))
}
func SkillCoverage(results []model.DrillResult) float64 {
	if len(results) == 0 {
		return 0
	}
	seen := map[string]bool{}
	for _, r := range results {
		seen[r.Drill] = true
	}
	return float64(len(seen))
}
func IsImproving(trend []float64) bool {
	if len(trend) < 2 {
		return false
	}
	return trend[len(trend)-1] > trend[0]
}
func IsDeclining(trend []float64) bool {
	if len(trend) < 2 {
		return false
	}
	return trend[len(trend)-1] < trend[0]
}
func TrendDelta(trend []float64) float64 {
	if len(trend) < 2 {
		return 0
	}
	return trend[len(trend)-1] - trend[0]
}
func SmoothTrend(trend []float64, window int) []float64 {
	if window < 1 {
		window = 1
	}
	out := make([]float64, len(trend))
	for i := range trend {
		start := i - window + 1
		if start < 0 {
			start = 0
		}
		total := 0.0
		for j := start; j <= i; j++ {
			total += trend[j]
		}
		out[i] = total / float64(i-start+1)
	}
	return out
}
func Percentile(values []float64, p float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	if p <= 0 {
		return sorted[0]
	}
	if p >= 100 {
		return sorted[len(sorted)-1]
	}
	idx := int(p / 100 * float64(len(sorted)-1))
	return sorted[idx]
}
func ResultValues(results []model.DrillResult) []float64 {
	out := make([]float64, 0, len(results))
	for _, r := range results {
		out = append(out, r.Score)
	}
	return out
}
func DrillScores(results []model.DrillResult, drill string) []float64 {
	out := []float64{}
	for _, r := range results {
		if r.Drill == drill {
			out = append(out, r.Score)
		}
	}
	return out
}
func AthletesWithResults(athletes []model.Athlete, results []model.DrillResult) []model.Athlete {
	seen := map[string]bool{}
	for _, r := range results {
		seen[r.AthleteID] = true
	}
	out := []model.Athlete{}
	for _, a := range athletes {
		if seen[a.ID] {
			out = append(out, a)
		}
	}
	return out
}
func ResultsForAthlete(results []model.DrillResult, id string) []model.DrillResult {
	out := []model.DrillResult{}
	for _, r := range results {
		if r.AthleteID == id {
			out = append(out, r)
		}
	}
	return out
}
