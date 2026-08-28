package analytics

import (
	"traininganalysis/internal/model"
	"traininganalysis/internal/training"
)

type CoachingNote struct {
	AthleteID string
	Category  string
	Priority  int
	Message   string
}

func GenerateNotes(results []model.DrillResult, benchmarks []Benchmark) []CoachingNote {
	notes := []CoachingNote{}
	for _, r := range results {
		status := ""
		for _, b := range benchmarks {
			if b.Drill == r.Drill {
				status = Compare(r, b)
			}
		}
		if status == "needs-work" {
			notes = append(notes, CoachingNote{AthleteID: r.AthleteID, Category: r.Drill, Priority: 3, Message: "repeat fundamentals"})
		} else if status == "watch" {
			notes = append(notes, CoachingNote{AthleteID: r.AthleteID, Category: r.Drill, Priority: 2, Message: "monitor next session"})
		} else if status == "on-track" {
			notes = append(notes, CoachingNote{AthleteID: r.AthleteID, Category: r.Drill, Priority: 1, Message: "increase challenge"})
		}
	}
	return notes
}
func GroupNotes(notes []CoachingNote) map[string][]CoachingNote {
	out := map[string][]CoachingNote{}
	for _, n := range notes {
		out[n.AthleteID] = append(out[n.AthleteID], n)
	}
	return out
}
func PriorityScore(notes []CoachingNote) int {
	total := 0
	for _, n := range notes {
		total += n.Priority
	}
	return total
}
func NextDrill(note CoachingNote) string {
	switch note.Category {
	case "shooting":
		return "movement-shot"
	case "passing":
		return "pressure-pass"
	case "defense":
		return "closeout"
	default:
		return "review"
	}
}
func SessionRecommendation(readiness, risk string) string {
	if risk == "critical" {
		return "rest"
	}
	if readiness == "recovery" {
		return "mobility"
	}
	if readiness == "developing" {
		return "technique"
	}
	return "competition"
}
func FilterNotes(notes []CoachingNote, minPriority int) []CoachingNote {
	out := []CoachingNote{}
	for _, n := range notes {
		if n.Priority >= minPriority {
			out = append(out, n)
		}
	}
	return out
}
func NoteCountByCategory(notes []CoachingNote) map[string]int {
	out := map[string]int{}
	for _, n := range notes {
		out[n.Category]++
	}
	return out
}
func AthleteReadiness(results []model.DrillResult) map[string]string {
	group := training.BuildEvaluations(results)
	out := map[string]string{}
	for id, e := range group {
		if e.Total >= 85 {
			out[id] = "ready"
		} else if e.Total >= 65 {
			out[id] = "developing"
		} else {
			out[id] = "recovery"
		}
	}
	return out
}
func SafeWorkload(sessions []model.TrainingSession, restDays int) int {
	load := Workload(sessions)
	if restDays > 0 {
		load -= restDays * 30
	}
	if load < 0 {
		return 0
	}
	return load
}
func ComparePeriods(a, b []model.DrillResult) map[string]float64 {
	old := training.BuildEvaluations(a)
	newer := training.BuildEvaluations(b)
	out := map[string]float64{}
	for id, e := range newer {
		if p := old[id]; p != nil {
			out[id] = e.Total - p.Total
		} else {
			out[id] = e.Total
		}
	}
	return out
}
func NormalizeResults(results []model.DrillResult) []model.DrillResult {
	out := make([]model.DrillResult, 0, len(results))
	for _, r := range results {
		r.Score = model.ClampScore(r.Score)
		if r.Attempts < 1 {
			r.Attempts = 1
		}
		out = append(out, r)
	}
	return out
}
func DrillVolume(results []model.DrillResult) map[string]int {
	out := map[string]int{}
	for _, r := range results {
		out[r.Drill] += r.Attempts
	}
	return out
}
func SessionIntensity(s model.TrainingSession) string {
	if s.Intensity >= 8 {
		return "high"
	}
	if s.Intensity >= 5 {
		return "moderate"
	}
	return "low"
}
func RecoveryRequired(sessions []model.TrainingSession) bool {
	for _, s := range sessions {
		if s.Intensity >= 9 || s.Duration >= 150 {
			return true
		}
	}
	return false
}
func BalanceScore(results []model.DrillResult) float64 {
	by := map[string][]model.DrillResult{}
	for _, r := range results {
		by[r.Drill] = append(by[r.Drill], r)
	}
	if len(by) == 0 {
		return 0
	}
	mins, maxs := 100.0, 0.0
	for _, list := range by {
		v := ScoreAverage(list)
		if v < mins {
			mins = v
		}
		if v > maxs {
			maxs = v
		}
	}
	return 100 - (maxs - mins)
}
