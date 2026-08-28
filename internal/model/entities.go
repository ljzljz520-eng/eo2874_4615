package model

type Athlete struct {
	ID, Name, AgeGroup, Position string
	BirthYear                    int
	Active                       bool
}
type TrainingSession struct {
	ID, TeamID, Date, Focus string
	Duration                int
	Intensity               int
}
type DrillResult struct {
	ID, SessionID, AthleteID, Drill string
	Score                           float64
	Attempts                        int
}
type TeamPlan struct {
	ID, TeamName, Season string
	Goals                []string
	WeeklyLoad           int
}
type Attendance struct {
	SessionID, AthleteID string
	Present              bool
	Note                 string
}
type Metric struct {
	Name  string
	Value float64
	Unit  string
}
