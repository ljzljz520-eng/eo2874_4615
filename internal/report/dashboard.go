package report

import (
	"fmt"
	"traininganalysis/internal/analytics"
	"traininganalysis/internal/model"
)

type Dashboard struct {
	Summary  analytics.Summary
	Workload int
	Risk     string
	Notes    []string
}

func BuildDashboard(athletes []model.Athlete, sessions []model.TrainingSession, capacity int) Dashboard {
	d := Dashboard{Summary: analytics.Summarize(athletes)}
	d.Workload = analytics.Workload(sessions)
	d.Risk = analytics.Risk(d.Workload, capacity)
	if d.Risk != "normal" {
		d.Notes = append(d.Notes, "adjust weekly load")
	}
	return d
}
func (d Dashboard) Lines() []string {
	return []string{fmt.Sprintf("total: %d", d.Summary.Total), fmt.Sprintf("workload: %d", d.Workload), "risk: " + d.Risk}
}
func (d Dashboard) Ready() bool { return d.Risk == "normal" && d.Summary.Total > 0 }
