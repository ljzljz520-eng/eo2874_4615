package model

import "sort"

func ActiveAthletes(in []Athlete) []Athlete {
	out := []Athlete{}
	for _, a := range in {
		if a.Active {
			out = append(out, a)
		}
	}
	return out
}
func FilterAge(in []Athlete, group string) []Athlete {
	out := []Athlete{}
	for _, a := range in {
		if a.AgeGroup == group {
			out = append(out, a)
		}
	}
	return out
}
func FilterPosition(in []Athlete, pos string) []Athlete {
	out := []Athlete{}
	for _, a := range in {
		if a.Position == pos {
			out = append(out, a)
		}
	}
	return out
}
func Names(in []Athlete) []string {
	out := []string{}
	for _, a := range in {
		out = append(out, a.Name)
	}
	return out
}
func IDs(in []Athlete) []string {
	out := []string{}
	for _, a := range in {
		out = append(out, a.ID)
	}
	return out
}
func UniqueGroups(in []Athlete) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, a := range in {
		if !seen[a.AgeGroup] {
			seen[a.AgeGroup] = true
			out = append(out, a.AgeGroup)
		}
	}
	return out
}
func UniquePositions(in []Athlete) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, a := range in {
		if !seen[a.Position] {
			seen[a.Position] = true
			out = append(out, a.Position)
		}
	}
	return out
}
func CopyAthletes(in []Athlete) []Athlete { out := make([]Athlete, len(in)); copy(out, in); return out }
func MergeAthletes(a, b []Athlete) []Athlete {
	out := CopyAthletes(a)
	seen := map[string]bool{}
	for _, v := range a {
		seen[v.ID] = true
	}
	for _, v := range b {
		if !seen[v.ID] {
			out = append(out, v)
			seen[v.ID] = true
		}
	}
	return out
}
func ReplaceAthlete(in []Athlete, v Athlete) []Athlete {
	out := CopyAthletes(in)
	for i, a := range out {
		if a.ID == v.ID {
			out[i] = v
			return out
		}
	}
	return append(out, v)
}
func RemoveAthlete(in []Athlete, id string) []Athlete {
	out := []Athlete{}
	for _, a := range in {
		if a.ID != id {
			out = append(out, a)
		}
	}
	return out
}
func SortAthletesByAge(in []Athlete) []Athlete {
	out := CopyAthletes(in)
	sort.Slice(out, func(i, j int) bool { return out[i].BirthYear < out[j].BirthYear })
	return out
}
func SortAthletesByGroup(in []Athlete) []Athlete {
	out := CopyAthletes(in)
	sort.Slice(out, func(i, j int) bool { return out[i].AgeGroup < out[j].AgeGroup })
	return out
}
func AgeCounts(in []Athlete) map[string]int {
	out := map[string]int{}
	for _, a := range in {
		out[a.AgeGroup]++
	}
	return out
}
func PositionCounts(in []Athlete) map[string]int {
	out := map[string]int{}
	for _, a := range in {
		out[a.Position]++
	}
	return out
}
func ValidateRoster(in []Athlete) bool {
	seen := map[string]bool{}
	for _, a := range in {
		if !a.Valid() || seen[a.ID] {
			return false
		}
		seen[a.ID] = true
	}
	return true
}
func FindAthlete(in []Athlete, id string) (Athlete, bool) {
	for _, a := range in {
		if a.ID == id {
			return a, true
		}
	}
	return Athlete{}, false
}
func ToggleActive(in []Athlete, id string) []Athlete {
	out := CopyAthletes(in)
	for i, a := range out {
		if a.ID == id {
			out[i].Active = !a.Active
		}
	}
	return out
}
func BirthYearRange(in []Athlete) (int, int) {
	if len(in) == 0 {
		return 0, 0
	}
	lo, hi := in[0].BirthYear, in[0].BirthYear
	for _, a := range in[1:] {
		if a.BirthYear < lo {
			lo = a.BirthYear
		}
		if a.BirthYear > hi {
			hi = a.BirthYear
		}
	}
	return lo, hi
}
func GroupActiveCount(in []Athlete, group string) int {
	n := 0
	for _, a := range in {
		if a.AgeGroup == group && a.Active {
			n++
		}
	}
	return n
}
func PositionActiveCount(in []Athlete, pos string) int {
	n := 0
	for _, a := range in {
		if a.Position == pos && a.Active {
			n++
		}
	}
	return n
}
