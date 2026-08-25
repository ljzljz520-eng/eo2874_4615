package schedule

import (
	"testing"
	"traininganalysis/internal/model"
)

func TestPlanner(t *testing.T) {
	p := NewPlanner()
	v := model.TeamPlan{ID: "p", TeamName: "T", Season: "2026", Goals: []string{"tech", "speed"}, WeeklyLoad: 100}
	p.Add(v)
	if len(p.ForSeason("2026")) != 1 {
		t.Fatal("season")
	}
	if p.NextFocus(v, 3) != "speed" {
		t.Fatal("focus")
	}
}
