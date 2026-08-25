package training

import (
	"fmt"
	"time"
	"traininganalysis/internal/model"
)

type SessionBook struct {
	sessions   map[string]model.TrainingSession
	attendance map[string][]model.Attendance
}

func NewSessionBook() *SessionBook {
	return &SessionBook{sessions: map[string]model.TrainingSession{}, attendance: map[string][]model.Attendance{}}
}
func (b *SessionBook) Schedule(team, focus string, duration, intensity int) model.TrainingSession {
	id := fmt.Sprintf("S-%d", time.Now().UnixNano())
	s := model.TrainingSession{ID: id, TeamID: team, Date: time.Now().Format("2006-01-02"), Focus: focus, Duration: duration, Intensity: intensity}
	if s.Valid() {
		b.sessions[id] = s
	}
	return s
}
func (b *SessionBook) RecordAttendance(a model.Attendance) {
	b.attendance[a.SessionID] = append(b.attendance[a.SessionID], a)
}
func (b *SessionBook) Session(id string) (model.TrainingSession, bool) {
	s, ok := b.sessions[id]
	return s, ok
}
func (b *SessionBook) AttendanceRate(id string) float64 {
	list := b.attendance[id]
	if len(list) == 0 {
		return 0
	}
	n := 0
	for _, a := range list {
		if a.Present {
			n++
		}
	}
	return float64(n) / float64(len(list)) * 100
}
