package storage

import "time"

type AuditEvent struct {
	Entity string
	Key    string
	Action string
	At     time.Time
}
type AuditLog struct{ events []AuditEvent }

func NewAuditLog() *AuditLog { return &AuditLog{events: []AuditEvent{}} }
func (a *AuditLog) Record(entity, key, action string) {
	if entity == "" || key == "" {
		return
	}
	a.events = append(a.events, AuditEvent{Entity: entity, Key: key, Action: action, At: time.Now()})
}
func (a *AuditLog) Events() []AuditEvent { return append([]AuditEvent(nil), a.events...) }
func (a *AuditLog) Count(entity string) int {
	n := 0
	for _, e := range a.events {
		if e.Entity == entity {
			n++
		}
	}
	return n
}
