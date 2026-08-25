package analytics

func Forecast(load int, weeks int) []int {
	if weeks < 0 {
		weeks = 0
	}
	out := make([]int, weeks)
	for i := range out {
		factor := 100 + i*5
		if i%3 == 0 {
			factor -= 10
		}
		out[i] = load * factor / 100
	}
	return out
}
func Risk(load, capacity int) string {
	if capacity <= 0 {
		return "critical"
	}
	ratio := float64(load) / float64(capacity)
	if ratio > 1.2 {
		return "critical"
	}
	if ratio > 0.9 {
		return "watch"
	}
	return "normal"
}
