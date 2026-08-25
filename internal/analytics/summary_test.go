package analytics

import (
	"testing"
	"traininganalysis/internal/model"
)

func TestSummary(t *testing.T) {
	s := Summarize([]model.Athlete{{Name: "Z"}})
	if s.Total != 1 {
		t.Fatal(s)
	}
}
func TestWorkflowSessionReview(t *testing.T) {
	if Readiness([]model.DrillResult{{Score: 90}}) != "ready" {
		t.Fatal("readiness")
	}
}
