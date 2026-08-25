package schedule

import (
	"sort"
	"traininganalysis/internal/model"
)

type Planner struct{ plans []model.TeamPlan }

func NewPlanner() *Planner { return &Planner{plans: []model.TeamPlan{}} }
func (p *Planner) Add(plan model.TeamPlan) {
	if plan.ID != "" {
		p.plans = append(p.plans, plan)
	}
}
func (p *Planner) ForSeason(season string) []model.TeamPlan {
	out := []model.TeamPlan{}
	for _, v := range p.plans {
		if v.Season == season {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].TeamName < out[j].TeamName })
	return out
}
func (p *Planner) NextFocus(plan model.TeamPlan, week int) string {
	if len(plan.Goals) == 0 {
		return "recovery"
	}
	return plan.Goals[week%len(plan.Goals)]
}
func (p *Planner) Load(plan model.TeamPlan, sessions int) int {
	if sessions < 0 {
		return 0
	}
	return plan.WeeklyLoad * sessions
}
