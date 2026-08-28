package metrics

import "traininganalysis/internal/model"

type Registry struct{ values map[string]model.Metric }

func New() *Registry { return &Registry{values: map[string]model.Metric{}} }
func (r *Registry) Set(name string, value float64, unit string) {
	if name != "" {
		r.values[name] = model.Metric{Name: name, Value: value, Unit: unit}
	}
}
func (r *Registry) Get(name string) (model.Metric, bool) { v, ok := r.values[name]; return v, ok }
func (r *Registry) Snapshot() []model.Metric {
	out := make([]model.Metric, 0, len(r.values))
	for _, v := range r.values {
		out = append(out, v)
	}
	return out
}
func (r *Registry) Merge(other *Registry) {
	for _, v := range other.values {
		r.Set(v.Name, v.Value, v.Unit)
	}
}
