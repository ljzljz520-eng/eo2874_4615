package report

import (
	"path/filepath"
	"testing"
	"traininganalysis/internal/model"
)

func TestReport(t *testing.T) {
	r := Build("x", []model.Athlete{{Name: "A"}}, nil)
	if r.Empty() {
		t.Fatal("empty")
	}
	p := filepath.Join(t.TempDir(), "a.csv")
	if e := ExportCSV(p, []model.Athlete{{ID: "1", Name: "A"}}); e != nil {
		t.Fatal(e)
	}
	if _, e := ImportCSV(p); e != nil {
		t.Fatal(e)
	}
}
