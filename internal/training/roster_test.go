package training

import (
	"os"
	"testing"
	"traininganalysis/internal/model"
)

func TestRosterFilterReset(t *testing.T) {
	if os.Getenv("RUN_BUG_REGRESSION") != "1" {
		t.Skip("regression scenario disabled in baseline run")
	}
	r := NewRoster()
	r.Add(model.Athlete{ID: "1", Name: "A", AgeGroup: "U16", BirthYear: 2008, Active: true})
	r.Add(model.Athlete{ID: "2", Name: "B", AgeGroup: "U18", BirthYear: 2006, Active: true})
	got := r.Filter("U16", 1, 10)
	if len(got) != 1 {
		t.Fatalf("u16=%d", len(got))
	}
	got = r.Filter("U18", 1, 10)
	if len(got) != 1 || got[0].AgeGroup != "U18" {
		t.Fatalf("stale roster: %#v", got)
	}
}
func TestWorkflowRosterAnalysis(t *testing.T) {
	r := NewRoster()
	for i := 0; i < 3; i++ {
		r.Add(model.Athlete{ID: string(rune('a' + i)), Name: "P", AgeGroup: "U16", BirthYear: 2008, Active: true})
	}
	if len(r.Filter("U16", 1, 2)) != 2 {
		t.Fatal("first page")
	}
}
