package report

import (
	"fmt"
	"sort"
	"strings"
	"traininganalysis/internal/model"
)

type Report struct {
	Title string
	Lines []string
}

func Build(title string, athletes []model.Athlete, results []model.DrillResult) Report {
	r := Report{Title: title}
	sort.Slice(athletes, func(i, j int) bool { return athletes[i].Name < athletes[j].Name })
	for _, a := range athletes {
		r.Lines = append(r.Lines, fmt.Sprintf("%s | %s | %s", a.Name, a.AgeGroup, a.Position))
	}
	if len(results) > 0 {
		r.Lines = append(r.Lines, fmt.Sprintf("results=%d", len(results)))
	}
	return r
}
func (r Report) Text() string { return r.Title + "\n" + strings.Join(r.Lines, "\n") }
func (r Report) Empty() bool  { return len(r.Lines) == 0 }
