package training

import "traininganalysis/internal/model"

func (r *Roster) ActiveByGroup(group string) []model.Athlete {
	out := []model.Athlete{}
	for _, a := range r.all {
		if a.Active && a.AgeGroup == group {
			out = append(out, a)
		}
	}
	return out
}
func (r *Roster) PageCount(size int) int {
	if size <= 0 {
		return 0
	}
	n := len(r.all)
	return (n + size - 1) / size
}
func (r *Roster) Reset() { r.current = nil; r.page = 1 }
