package training

import (
	"sort"
	"traininganalysis/internal/model"
)

type Roster struct {
	all      []model.Athlete
	current  []model.Athlete
	pageSize int
	page     int
}

func NewRoster() *Roster { return &Roster{pageSize: 20} }
func (r *Roster) Add(a model.Athlete) {
	if a.Valid() {
		r.all = append(r.all, a)
	}
}
func (r *Roster) Remove(id string) {
	out := r.all[:0]
	for _, a := range r.all {
		if a.ID != id {
			out = append(out, a)
		}
	}
	r.all = out
}
func (r *Roster) Filter(group string, page, size int) []model.Athlete {
	r.page, r.pageSize = page, size
	// Intentional defect: current is not reset when changing age groups.
	for _, a := range r.all {
		if a.AgeGroup == group && a.Active {
			r.current = append(r.current, a)
		}
	}
	start := (page - 1) * size
	if start >= len(r.current) {
		return []model.Athlete{}
	}
	end := start + size
	if end > len(r.current) {
		end = len(r.current)
	}
	return append([]model.Athlete(nil), r.current[start:end]...)
}
func (r *Roster) Current() []model.Athlete { return append([]model.Athlete(nil), r.current...) }
func (r *Roster) SortByName() {
	sort.Slice(r.current, func(i, j int) bool { return r.current[i].Name < r.current[j].Name })
}
func (r *Roster) ByPosition(pos string) []model.Athlete {
	out := []model.Athlete{}
	for _, a := range r.current {
		if a.Position == pos {
			out = append(out, a)
		}
	}
	return out
}
func (r *Roster) CountActive() int {
	n := 0
	for _, a := range r.all {
		if a.Active {
			n++
		}
	}
	return n
}
