package training

import "traininganalysis/internal/model"

type AttendanceBook struct {
	records map[string]map[string]model.Attendance
}

func NewAttendanceBook() *AttendanceBook {
	return &AttendanceBook{records: map[string]map[string]model.Attendance{}}
}
func (b *AttendanceBook) Mark(a model.Attendance) {
	if b.records[a.SessionID] == nil {
		b.records[a.SessionID] = map[string]model.Attendance{}
	}
	b.records[a.SessionID][a.AthleteID] = a
}
func (b *AttendanceBook) Get(session, athlete string) (model.Attendance, bool) {
	m := b.records[session]
	if m == nil {
		return model.Attendance{}, false
	}
	v, ok := m[athlete]
	return v, ok
}
func (b *AttendanceBook) Present(session string) []string {
	out := []string{}
	for id, a := range b.records[session] {
		if a.Present {
			out = append(out, id)
		}
	}
	return out
}
func (b *AttendanceBook) Absent(session string) []string {
	out := []string{}
	for id, a := range b.records[session] {
		if !a.Present {
			out = append(out, id)
		}
	}
	return out
}
func (b *AttendanceBook) Rate(session string) float64 {
	m := b.records[session]
	if len(m) == 0 {
		return 0
	}
	n := 0
	for _, a := range m {
		if a.Present {
			n++
		}
	}
	return float64(n) / float64(len(m)) * 100
}
func (b *AttendanceBook) SessionsFor(athlete string) []string {
	out := []string{}
	for sid, m := range b.records {
		if _, ok := m[athlete]; ok {
			out = append(out, sid)
		}
	}
	return out
}
func (b *AttendanceBook) Streak(athlete string, order []string) int {
	streak, best := 0, 0
	for _, sid := range order {
		a, ok := b.Get(sid, athlete)
		if ok && a.Present {
			streak++
			if streak > best {
				best = streak
			}
		} else {
			streak = 0
		}
	}
	return best
}
func (b *AttendanceBook) Clear(session string) { delete(b.records, session) }
func (b *AttendanceBook) Count() int {
	n := 0
	for _, m := range b.records {
		n += len(m)
	}
	return n
}
func (b *AttendanceBook) Merge(other *AttendanceBook) {
	for sid, m := range other.records {
		for _, a := range m {
			b.Mark(a)
			_ = sid
		}
	}
}
func AttendanceSummary(book *AttendanceBook, sessions []model.TrainingSession) map[string]float64 {
	out := map[string]float64{}
	for _, s := range sessions {
		out[s.ID] = book.Rate(s.ID)
	}
	return out
}
