package storage

import (
	"path/filepath"
	"testing"
	"traininganalysis/internal/model"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	a := model.Athlete{ID: "a", Name: "Persist", AgeGroup: "U16", BirthYear: 2008, Active: true}
	if e = s.SaveAthlete(a); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.LoadAthlete("a")
	if e != nil || got.Name != "Persist" {
		t.Fatalf("%v %#v", e, got)
	}
}
