package schedule

import "traininganalysis/internal/model"

type Phase struct {
	Name      string
	Weeks     int
	Intensity int
	Focus     string
}

func DefaultPhases() []Phase {
	return []Phase{{Name: "base", Weeks: 4, Intensity: 5, Focus: "technique"}, {Name: "build", Weeks: 5, Intensity: 7, Focus: "capacity"}, {Name: "peak", Weeks: 3, Intensity: 9, Focus: "competition"}, {Name: "taper", Weeks: 2, Intensity: 4, Focus: "recovery"}}
}
func PhaseForWeek(phases []Phase, week int) (Phase, bool) {
	if week < 1 {
		return Phase{}, false
	}
	offset := 0
	for _, p := range phases {
		offset += p.Weeks
		if week <= offset {
			return p, true
		}
	}
	return Phase{}, false
}
func SessionsForPhase(p Phase, team string) []model.TrainingSession {
	out := []model.TrainingSession{}
	for i := 0; i < p.Weeks; i++ {
		d := 90
		if p.Intensity >= 8 {
			d = 75
		}
		if p.Name == "taper" {
			d = 60
		}
		out = append(out, model.TrainingSession{ID: p.Name + "-" + string(rune(i)), TeamID: team, Focus: p.Focus, Duration: d, Intensity: p.Intensity})
	}
	return out
}
func PhaseLoad(p Phase) int {
	minutes := 90
	if p.Intensity >= 8 {
		minutes = 75
	}
	if p.Name == "taper" {
		minutes = 60
	}
	return p.Weeks * minutes * p.Intensity
}
func TotalPhaseLoad(phases []Phase) int {
	n := 0
	for _, p := range phases {
		n += PhaseLoad(p)
	}
	return n
}
func PeakPhase(phases []Phase) Phase {
	out := Phase{}
	for _, p := range phases {
		if p.Intensity > out.Intensity {
			out = p
		}
	}
	return out
}
func ValidatePhases(phases []Phase) bool {
	if len(phases) == 0 {
		return false
	}
	for _, p := range phases {
		if p.Weeks <= 0 || p.Intensity < 1 || p.Intensity > 10 {
			return false
		}
	}
	return true
}
