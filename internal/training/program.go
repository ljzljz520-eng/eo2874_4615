package training

import "traininganalysis/internal/model"

type Program struct {
	Name   string
	Blocks []ProgramBlock
}
type ProgramBlock struct {
	Name      string
	Goal      string
	Drills    []string
	Minutes   int
	Intensity int
}

func NewProgram(name string) *Program { return &Program{Name: name, Blocks: []ProgramBlock{}} }
func (p *Program) AddBlock(b ProgramBlock) {
	if b.Name == "" || b.Minutes <= 0 {
		return
	}
	if b.Intensity < 1 {
		b.Intensity = 1
	}
	if b.Intensity > 10 {
		b.Intensity = 10
	}
	p.Blocks = append(p.Blocks, b)
}
func (p *Program) TotalMinutes() int {
	n := 0
	for _, b := range p.Blocks {
		n += b.Minutes
	}
	return n
}
func (p *Program) AverageIntensity() float64 {
	if len(p.Blocks) == 0 {
		return 0
	}
	n := 0
	for _, b := range p.Blocks {
		n += b.Intensity
	}
	return float64(n) / float64(len(p.Blocks))
}
func (p *Program) Goals() []string {
	out := []string{}
	seen := map[string]bool{}
	for _, b := range p.Blocks {
		if b.Goal != "" && !seen[b.Goal] {
			seen[b.Goal] = true
			out = append(out, b.Goal)
		}
	}
	return out
}
func (p *Program) DrillCount() int {
	n := 0
	for _, b := range p.Blocks {
		n += len(b.Drills)
	}
	return n
}
func (p *Program) Validate() bool {
	if p.Name == "" || len(p.Blocks) == 0 {
		return false
	}
	for _, b := range p.Blocks {
		if b.Minutes <= 0 || b.Intensity < 1 {
			return false
		}
	}
	return true
}
func (p *Program) Block(name string) (ProgramBlock, bool) {
	for _, b := range p.Blocks {
		if b.Name == name {
			return b, true
		}
	}
	return ProgramBlock{}, false
}
func (p *Program) Replace(name string, b ProgramBlock) bool {
	for i, v := range p.Blocks {
		if v.Name == name {
			p.Blocks[i] = b
			return true
		}
	}
	return false
}
func (p *Program) Remove(name string) bool {
	for i, b := range p.Blocks {
		if b.Name == name {
			p.Blocks = append(p.Blocks[:i], p.Blocks[i+1:]...)
			return true
		}
	}
	return false
}
func BuildDefaultProgram() *Program {
	p := NewProgram("foundation")
	p.AddBlock(ProgramBlock{Name: "warmup", Goal: "mobility", Drills: []string{"jog", "dynamic"}, Minutes: 15, Intensity: 3})
	p.AddBlock(ProgramBlock{Name: "skill", Goal: "technique", Drills: []string{"passing", "shooting"}, Minutes: 35, Intensity: 6})
	p.AddBlock(ProgramBlock{Name: "game", Goal: "decision", Drills: []string{"small-sided"}, Minutes: 30, Intensity: 8})
	return p
}
func (p *Program) ToSession(team string) model.TrainingSession {
	return model.TrainingSession{ID: p.Name, TeamID: team, Focus: p.Name, Duration: p.TotalMinutes(), Intensity: int(p.AverageIntensity() + 0.5)}
}
func MergePrograms(a, b *Program) *Program {
	out := NewProgram(a.Name + "+" + b.Name)
	for _, v := range a.Blocks {
		out.AddBlock(v)
	}
	for _, v := range b.Blocks {
		if _, ok := out.Block(v.Name); !ok {
			out.AddBlock(v)
		}
	}
	return out
}
func ScaleProgram(p *Program, factor float64) *Program {
	out := NewProgram(p.Name)
	for _, b := range p.Blocks {
		b.Minutes = int(float64(b.Minutes) * factor)
		if b.Minutes < 1 {
			b.Minutes = 1
		}
		out.AddBlock(b)
	}
	return out
}
func (p *Program) IntensityBand() string {
	v := p.AverageIntensity()
	if v >= 8 {
		return "high"
	}
	if v >= 5 {
		return "moderate"
	}
	return "low"
}
