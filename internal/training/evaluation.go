package training

import (
	"sort"
	"traininganalysis/internal/model"
)

type Evaluation struct {
	AthleteID string
	Scores    map[string]float64
	Total     float64
	Grade     string
}

func NewEvaluation(id string) *Evaluation {
	return &Evaluation{AthleteID: id, Scores: map[string]float64{}}
}
func (e *Evaluation) Add(drill string, score float64) {
	if drill == "" {
		return
	}
	e.Scores[drill] = model.ClampScore(score)
	e.recalculate()
}
func (e *Evaluation) recalculate() {
	e.Total = 0
	for _, v := range e.Scores {
		e.Total += v
	}
	if len(e.Scores) > 0 {
		e.Total /= float64(len(e.Scores))
	}
	switch {
	case e.Total >= 90:
		e.Grade = "A"
	case e.Total >= 80:
		e.Grade = "B"
	case e.Total >= 70:
		e.Grade = "C"
	default:
		e.Grade = "D"
	}
}
func (e *Evaluation) Passed(threshold float64) bool { return e.Total >= threshold }
func RankEvaluations(items []*Evaluation) []*Evaluation {
	out := append([]*Evaluation(nil), items...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Total > out[j].Total })
	return out
}
func BuildEvaluations(results []model.DrillResult) map[string]*Evaluation {
	out := map[string]*Evaluation{}
	for _, r := range results {
		e := out[r.AthleteID]
		if e == nil {
			e = NewEvaluation(r.AthleteID)
			out[r.AthleteID] = e
		}
		e.Add(r.Drill, r.Score)
	}
	return out
}
func SelectTop(items []*Evaluation, n int) []*Evaluation {
	if n < 0 {
		n = 0
	}
	if n > len(items) {
		n = len(items)
	}
	ranked := RankEvaluations(items)
	return ranked[:n]
}
func Distribution(items []*Evaluation) map[string]int {
	out := map[string]int{}
	for _, e := range items {
		out[e.Grade]++
	}
	return out
}
