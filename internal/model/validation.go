package model

import "strings"

func (a Athlete) Valid() bool {
	return a.ID != "" && a.Name != "" && a.AgeGroup != "" && a.BirthYear > 1900
}
func (s TrainingSession) Valid() bool {
	return s.ID != "" && s.TeamID != "" && s.Duration > 0 && s.Intensity >= 1 && s.Intensity <= 10
}
func NormalizeName(v string) string { return strings.TrimSpace(strings.Title(strings.ToLower(v))) }
func ClampScore(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return 100
	}
	return v
}
