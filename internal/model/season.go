package model

type Season struct {
	Name   string
	Start  string
	End    string
	Teams  []string
	Status string
}

func NewSeason(name, start, end string) *Season {
	return &Season{Name: name, Start: start, End: end, Status: "planned"}
}
func (s *Season) AddTeam(team string) {
	if team != "" {
		for _, v := range s.Teams {
			if v == team {
				return
			}
		}
		s.Teams = append(s.Teams, team)
	}
}
func (s *Season) RemoveTeam(team string) bool {
	for i, v := range s.Teams {
		if v == team {
			s.Teams = append(s.Teams[:i], s.Teams[i+1:]...)
			return true
		}
	}
	return false
}
func (s *Season) HasTeam(team string) bool {
	for _, v := range s.Teams {
		if v == team {
			return true
		}
	}
	return false
}
func (s *Season) TeamCount() int { return len(s.Teams) }
func (s *Season) Activate() bool {
	if s.Name == "" || len(s.Teams) == 0 {
		return false
	}
	s.Status = "active"
	return true
}
func (s *Season) Complete()        { s.Status = "complete" }
func (s *Season) Cancel()          { s.Status = "cancelled" }
func (s *Season) IsActive() bool   { return s.Status == "active" }
func (s *Season) IsFinished() bool { return s.Status == "complete" || s.Status == "cancelled" }
func (s *Season) Validate() bool {
	if s.Name == "" || s.Start == "" || s.End == "" {
		return false
	}
	switch s.Status {
	case "planned", "active", "complete", "cancelled":
		return true
	default:
		return false
	}
}
func (s *Season) Clone() *Season {
	out := *s
	out.Teams = append([]string(nil), s.Teams...)
	return &out
}
func MergeSeasons(a, b *Season) *Season {
	out := a.Clone()
	for _, team := range b.Teams {
		out.AddTeam(team)
	}
	if out.End < b.End {
		out.End = b.End
	}
	return out
}
func SeasonStatus(s *Season) string {
	if s == nil {
		return "unknown"
	}
	return s.Status
}
func TeamNames(s *Season) []string {
	if s == nil {
		return nil
	}
	return append([]string(nil), s.Teams...)
}
func HasAnyTeam(s *Season, teams []string) bool {
	if s == nil {
		return false
	}
	for _, team := range teams {
		if s.HasTeam(team) {
			return true
		}
	}
	return false
}
func AllTeams(s *Season, teams []string) bool {
	if s == nil {
		return false
	}
	for _, team := range teams {
		if !s.HasTeam(team) {
			return false
		}
	}
	return true
}
func TeamOverlap(a, b *Season) []string {
	out := []string{}
	if a == nil || b == nil {
		return out
	}
	for _, team := range a.Teams {
		if b.HasTeam(team) {
			out = append(out, team)
		}
	}
	return out
}
func TeamDifference(a, b *Season) []string {
	out := []string{}
	if a == nil {
		return out
	}
	for _, team := range a.Teams {
		if b == nil || !b.HasTeam(team) {
			out = append(out, team)
		}
	}
	return out
}
func SeasonPhase(s *Season) string {
	if s == nil {
		return "none"
	}
	switch s.Status {
	case "planned":
		return "setup"
	case "active":
		return "delivery"
	case "complete":
		return "review"
	default:
		return "closed"
	}
}
func CanSchedule(s *Season) bool { return s != nil && s.Status == "active" && len(s.Teams) > 0 }
