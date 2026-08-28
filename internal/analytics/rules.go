package analytics

import "traininganalysis/internal/model"

type Rule struct {
	Name      string
	Threshold float64
	Message   string
}

func DefaultRules() []Rule {
	return []Rule{{Name: "excellent", Threshold: 90, Message: "promote challenge"}, {Name: "strong", Threshold: 80, Message: "maintain load"}, {Name: "developing", Threshold: 65, Message: "add repetition"}, {Name: "recovery", Threshold: 0, Message: "reduce load"}}
}
func ApplyRule(score float64, rules []Rule) Rule {
	best := Rule{}
	for _, r := range rules {
		if score >= r.Threshold && r.Threshold >= best.Threshold {
			best = r
		}
	}
	return best
}
func RuleMessages(results []model.DrillResult, rules []Rule) map[string]string {
	out := map[string]string{}
	for _, r := range results {
		out[r.AthleteID+":"+r.Drill] = ApplyRule(r.Score, rules).Message
	}
	return out
}
func Thresholds(rules []Rule) []float64 {
	out := []float64{}
	for _, r := range rules {
		out = append(out, r.Threshold)
	}
	return out
}
func ValidateRules(rules []Rule) bool {
	if len(rules) == 0 {
		return false
	}
	for i, r := range rules {
		if r.Name == "" || r.Message == "" || r.Threshold < 0 || r.Threshold > 100 {
			return false
		}
		if i > 0 && r.Threshold > rules[i-1].Threshold {
			return false
		}
	}
	return true
}
func MissingDrills(results []model.DrillResult, required []string) map[string][]string {
	seen := map[string]map[string]bool{}
	for _, r := range results {
		if seen[r.AthleteID] == nil {
			seen[r.AthleteID] = map[string]bool{}
		}
		seen[r.AthleteID][r.Drill] = true
	}
	out := map[string][]string{}
	for id, drills := range seen {
		for _, req := range required {
			if !drills[req] {
				out[id] = append(out[id], req)
			}
		}
	}
	return out
}
func ScoreBuckets(results []model.DrillResult) map[string]int {
	out := map[string]int{}
	for _, r := range results {
		b := "low"
		if r.Score >= 80 {
			b = "high"
		} else if r.Score >= 60 {
			b = "mid"
		}
		out[b]++
	}
	return out
}
