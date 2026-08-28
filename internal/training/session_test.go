package training

import (
	"testing"
	"traininganalysis/internal/model"
)

func TestSessionBook(t *testing.T) {
	b := NewSessionBook()
	s := b.Schedule("T", "speed", 60, 7)
	b.RecordAttendance(model.Attendance{SessionID: s.ID, AthleteID: "a", Present: true})
	if b.AttendanceRate(s.ID) != 100 {
		t.Fatal("attendance")
	}
}
func TestWorkflowPlanSession(t *testing.T) {
	b := NewSessionBook()
	if s := b.Schedule("T", "skills", 90, 5); s.ID == "" {
		t.Fatal("id")
	}
}
