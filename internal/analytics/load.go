package analytics

import "traininganalysis/internal/model"

func Workload(sessions []model.TrainingSession) int {
	total := 0
	for _, s := range sessions {
		factor := s.Intensity
		if factor < 1 {
			factor = 1
		}
		total += s.Duration * factor
	}
	return total
}
func Readiness(results []model.DrillResult) string {
	avg := ScoreAverage(results)
	switch {
	case avg >= 85:
		return "ready"
	case avg >= 65:
		return "developing"
	default:
		return "recovery"
	}
}
func PositionBreakdown(athletes []model.Athlete) map[string]int {
	out := map[string]int{}
	for _, a := range athletes {
		out[a.Position]++
	}
	return out
}
