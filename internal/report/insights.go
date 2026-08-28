package report

import (
	"fmt"
	"os"
	"sort"
	"traininganalysis/internal/analytics"
	"traininganalysis/internal/model"
)

type Insight struct {
	Heading  string
	Detail   string
	Severity string
}

func BuildInsights(athletes []model.Athlete, sessions []model.TrainingSession, results []model.DrillResult) []Insight {
	out := []Insight{}
	if len(athletes) == 0 {
		out = append(out, Insight{Heading: "Roster", Detail: "No active athletes", Severity: "high"})
	}
	if analytics.RecoveryRequired(sessions) {
		out = append(out, Insight{Heading: "Recovery", Detail: "High intensity block detected", Severity: "medium"})
	}
	if len(results) > 0 {
		readiness := analytics.Readiness(results)
		out = append(out, Insight{Heading: "Readiness", Detail: readiness, Severity: severity(readiness)})
	}
	return out
}
func severity(v string) string {
	switch v {
	case "ready":
		return "low"
	case "developing":
		return "medium"
	default:
		return "high"
	}
}
func SortInsights(in []Insight) []Insight {
	out := append([]Insight(nil), in...)
	rank := map[string]int{"high": 3, "medium": 2, "low": 1}
	sort.SliceStable(out, func(i, j int) bool { return rank[out[i].Severity] > rank[out[j].Severity] })
	return out
}
func RenderInsights(in []Insight) []string {
	out := []string{}
	for _, v := range SortInsights(in) {
		out = append(out, fmt.Sprintf("[%s] %s: %s", v.Severity, v.Heading, v.Detail))
	}
	return out
}
func MergeInsights(a, b []Insight) []Insight {
	out := append([]Insight(nil), a...)
	out = append(out, b...)
	return SortInsights(out)
}
func HasSeverity(in []Insight, want string) bool {
	for _, v := range in {
		if v.Severity == want {
			return true
		}
	}
	return false
}
func SummaryLine(in []Insight) string {
	counts := map[string]int{}
	for _, v := range in {
		counts[v.Severity]++
	}
	return fmt.Sprintf("high=%d medium=%d low=%d", counts["high"], counts["medium"], counts["low"])
}
func AthleteLines(athletes []model.Athlete) []string {
	out := []string{}
	for _, a := range athletes {
		status := "inactive"
		if a.Active {
			status = "active"
		}
		out = append(out, fmt.Sprintf("%s (%s) %s", a.Name, a.AgeGroup, status))
	}
	return out
}
func SessionLines(sessions []model.TrainingSession) []string {
	out := []string{}
	for _, s := range sessions {
		out = append(out, fmt.Sprintf("%s %s %dmin intensity-%d", s.Date, s.Focus, s.Duration, s.Intensity))
	}
	return out
}
func ResultLines(results []model.DrillResult) []string {
	out := []string{}
	for _, r := range results {
		out = append(out, fmt.Sprintf("%s %s %.1f/%d", r.AthleteID, r.Drill, r.Score, r.Attempts))
	}
	return out
}
func ExportText(path string, lines []string) error { return writeLines(path, lines) }
func writeLines(path string, lines []string) error {
	f, e := os.Create(path)
	if e != nil {
		return e
	}
	defer f.Close()
	for _, line := range lines {
		if _, e = f.WriteString(line + "\n"); e != nil {
			return e
		}
	}
	return nil
}
