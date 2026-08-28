package schedule

import "traininganalysis/internal/model"

func SessionDays(sessions []model.TrainingSession) map[string]int {
	out := map[string]int{}
	for _, s := range sessions {
		out[s.Date]++
	}
	return out
}
func FocusCounts(sessions []model.TrainingSession) map[string]int {
	out := map[string]int{}
	for _, s := range sessions {
		out[s.Focus]++
	}
	return out
}
func IntensityCounts(sessions []model.TrainingSession) map[int]int {
	out := map[int]int{}
	for _, s := range sessions {
		out[s.Intensity]++
	}
	return out
}
func DurationRange(sessions []model.TrainingSession) (int, int) {
	if len(sessions) == 0 {
		return 0, 0
	}
	lo, hi := sessions[0].Duration, sessions[0].Duration
	for _, s := range sessions[1:] {
		if s.Duration < lo {
			lo = s.Duration
		}
		if s.Duration > hi {
			hi = s.Duration
		}
	}
	return lo, hi
}
func LongestSession(sessions []model.TrainingSession) (model.TrainingSession, bool) {
	if len(sessions) == 0 {
		return model.TrainingSession{}, false
	}
	best := sessions[0]
	for _, s := range sessions[1:] {
		if s.Duration > best.Duration {
			best = s
		}
	}
	return best, true
}
func HighIntensity(sessions []model.TrainingSession, min int) []model.TrainingSession {
	out := []model.TrainingSession{}
	for _, s := range sessions {
		if s.Intensity >= min {
			out = append(out, s)
		}
	}
	return out
}
func LowIntensity(sessions []model.TrainingSession, max int) []model.TrainingSession {
	out := []model.TrainingSession{}
	for _, s := range sessions {
		if s.Intensity <= max {
			out = append(out, s)
		}
	}
	return out
}
func SessionsForTeam(sessions []model.TrainingSession, team string) []model.TrainingSession {
	out := []model.TrainingSession{}
	for _, s := range sessions {
		if s.TeamID == team {
			out = append(out, s)
		}
	}
	return out
}
func SessionsForFocus(sessions []model.TrainingSession, focus string) []model.TrainingSession {
	out := []model.TrainingSession{}
	for _, s := range sessions {
		if s.Focus == focus {
			out = append(out, s)
		}
	}
	return out
}
func ValidSessions(sessions []model.TrainingSession) bool {
	for _, s := range sessions {
		if !s.Valid() {
			return false
		}
	}
	return true
}
func CopySessions(sessions []model.TrainingSession) []model.TrainingSession {
	out := make([]model.TrainingSession, len(sessions))
	copy(out, sessions)
	return out
}
