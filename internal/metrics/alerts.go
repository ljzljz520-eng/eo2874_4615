package metrics

import "traininganalysis/internal/model"

type Alert struct {
	Name  string
	Value float64
	Limit float64
	Level string
}

func Evaluate(name string, value, limit float64) Alert {
	level := "ok"
	if value > limit*1.2 {
		level = "critical"
	} else if value > limit {
		level = "warning"
	}
	return Alert{Name: name, Value: value, Limit: limit, Level: level}
}
func CriticalAlerts(alerts []Alert) []Alert {
	out := []Alert{}
	for _, a := range alerts {
		if a.Level == "critical" {
			out = append(out, a)
		}
	}
	return out
}
func WarningAlerts(alerts []Alert) []Alert {
	out := []Alert{}
	for _, a := range alerts {
		if a.Level == "warning" {
			out = append(out, a)
		}
	}
	return out
}
func AlertMap(alerts []Alert) map[string]Alert {
	out := map[string]Alert{}
	for _, a := range alerts {
		out[a.Name] = a
	}
	return out
}
func AlertValues(alerts []Alert) []model.Metric {
	out := []model.Metric{}
	for _, a := range alerts {
		out = append(out, model.Metric{Name: a.Name, Value: a.Value, Unit: a.Level})
	}
	return out
}
func WorstAlert(alerts []Alert) Alert {
	best := Alert{}
	rank := map[string]int{"ok": 0, "warning": 1, "critical": 2}
	for _, a := range alerts {
		if rank[a.Level] > rank[best.Level] {
			best = a
		}
	}
	return best
}
func Resolve(alert Alert, value float64) Alert {
	alert.Value = value
	if value <= alert.Limit {
		alert.Level = "ok"
	} else if value <= alert.Limit*1.2 {
		alert.Level = "warning"
	} else {
		alert.Level = "critical"
	}
	return alert
}
func AlertCount(alerts []Alert, level string) int {
	n := 0
	for _, a := range alerts {
		if a.Level == level {
			n++
		}
	}
	return n
}
func Healthy(alerts []Alert) bool {
	for _, a := range alerts {
		if a.Level == "critical" {
			return false
		}
	}
	return true
}
func Levels(alerts []Alert) []string {
	out := []string{}
	for _, a := range alerts {
		out = append(out, a.Level)
	}
	return out
}
func Names(alerts []Alert) []string {
	out := []string{}
	for _, a := range alerts {
		out = append(out, a.Name)
	}
	return out
}
func Limits(alerts []Alert) []float64 {
	out := []float64{}
	for _, a := range alerts {
		out = append(out, a.Limit)
	}
	return out
}
func Values(alerts []Alert) []float64 {
	out := []float64{}
	for _, a := range alerts {
		out = append(out, a.Value)
	}
	return out
}
func OverLimit(alerts []Alert) []Alert {
	out := []Alert{}
	for _, a := range alerts {
		if a.Value > a.Limit {
			out = append(out, a)
		}
	}
	return out
}
func UnderLimit(alerts []Alert) []Alert {
	out := []Alert{}
	for _, a := range alerts {
		if a.Value <= a.Limit {
			out = append(out, a)
		}
	}
	return out
}
func Escalate(alert Alert) Alert {
	if alert.Level == "warning" {
		alert.Level = "critical"
	}
	return alert
}
func Downgrade(alert Alert) Alert {
	if alert.Level == "critical" {
		alert.Level = "warning"
	} else if alert.Level == "warning" {
		alert.Level = "ok"
	}
	return alert
}
func Combine(a, b Alert) Alert {
	if a.Name == "" {
		return b
	}
	if b.Name == "" {
		return a
	}
	if a.Level == "critical" || b.Level == "critical" {
		a.Level = "critical"
	} else if a.Level == "warning" || b.Level == "warning" {
		a.Level = "warning"
	}
	if b.Value > a.Value {
		a.Value = b.Value
	}
	return a
}
