package schedule

import "traininganalysis/internal/model"

type Cycle struct {
	Name  string
	Weeks []Week
}
type Week struct {
	Number   int
	Sessions []model.TrainingSession
	Recovery bool
}

func BuildCycle(name string, weeks int, team string) *Cycle {
	c := &Cycle{Name: name}
	if weeks < 0 {
		weeks = 0
	}
	for i := 1; i <= weeks; i++ {
		w := Week{Number: i}
		if i%4 == 0 {
			w.Recovery = true
		}
		w.Sessions = append(w.Sessions, model.TrainingSession{ID: team + "-" + string(rune(i)), TeamID: team, Focus: "technical", Duration: 90, Intensity: 6})
		c.Weeks = append(c.Weeks, w)
	}
	return c
}
func (c *Cycle) TotalMinutes() int {
	n := 0
	for _, w := range c.Weeks {
		for _, s := range w.Sessions {
			n += s.Duration
		}
	}
	return n
}
func (c *Cycle) RecoveryWeeks() []int {
	out := []int{}
	for _, w := range c.Weeks {
		if w.Recovery {
			out = append(out, w.Number)
		}
	}
	return out
}
func (c *Cycle) SessionCount() int {
	n := 0
	for _, w := range c.Weeks {
		n += len(w.Sessions)
	}
	return n
}
