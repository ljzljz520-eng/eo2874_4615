package training

import (
	"testing"
	"traininganalysis/internal/model"
)

func TestEvaluation(t *testing.T) {
	e := NewEvaluation("a")
	e.Add("shot", 90)
	e.Add("pass", 80)
	if e.Grade != "B" {
		t.Fatal(e.Grade)
	}
}
func TestBuildEvaluations(t *testing.T) {
	m := BuildEvaluations([]model.DrillResult{{AthleteID: "a", Drill: "x", Score: 90}})
	if len(m) != 1 {
		t.Fatal(len(m))
	}
}
