package model

import "strings"

type DrillCatalog struct{ groups map[string][]string }

func NewDrillCatalog() *DrillCatalog {
	return &DrillCatalog{groups: map[string][]string{"technical": {"passing", "shooting", "dribbling"}, "physical": {"speed", "agility", "strength"}, "tactical": {"spacing", "pressing", "transition"}}}
}
func (c *DrillCatalog) Add(group, drill string) {
	if group == "" || drill == "" {
		return
	}
	c.groups[group] = append(c.groups[group], drill)
}
func (c *DrillCatalog) Drills(group string) []string {
	return append([]string(nil), c.groups[group]...)
}
func (c *DrillCatalog) Groups() []string {
	out := []string{}
	for g := range c.groups {
		out = append(out, g)
	}
	return out
}
func (c *DrillCatalog) Has(group, drill string) bool {
	for _, v := range c.groups[group] {
		if v == drill {
			return true
		}
	}
	return false
}
func (c *DrillCatalog) Remove(group, drill string) bool {
	list := c.groups[group]
	for i, v := range list {
		if v == drill {
			c.groups[group] = append(list[:i], list[i+1:]...)
			return true
		}
	}
	return false
}
func (c *DrillCatalog) Count(group string) int { return len(c.groups[group]) }
func (c *DrillCatalog) Total() int {
	n := 0
	for _, v := range c.groups {
		n += len(v)
	}
	return n
}
func (c *DrillCatalog) Clone() *DrillCatalog {
	out := NewDrillCatalog()
	out.groups = map[string][]string{}
	for g, v := range c.groups {
		out.groups[g] = append([]string(nil), v...)
	}
	return out
}
func (c *DrillCatalog) Search(term string) map[string][]string {
	out := map[string][]string{}
	for g, v := range c.groups {
		for _, d := range v {
			if strings.Contains(d, term) {
				out[g] = append(out[g], d)
			}
		}
	}
	return out
}
func (c *DrillCatalog) Flatten() []string {
	out := []string{}
	for _, v := range c.groups {
		out = append(out, v...)
	}
	return out
}
func (c *DrillCatalog) Validate() bool {
	for g, v := range c.groups {
		if g == "" || len(v) == 0 {
			return false
		}
		for _, d := range v {
			if d == "" {
				return false
			}
		}
	}
	return true
}
