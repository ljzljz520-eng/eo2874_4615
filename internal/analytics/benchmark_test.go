package analytics

import (
	"testing"
	"traininganalysis/internal/model"
)

func TestBenchmark(t *testing.T) {
	if Compare(model.DrillResult{Score: 80}, Benchmark{Target: 85, Tolerance: 3}) != "watch" {
		t.Fatal("benchmark")
	}
}
