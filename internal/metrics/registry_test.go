package metrics

import "testing"

func TestRegistry(t *testing.T) {
	r := New()
	r.Set("load", 10, "au")
	if _, ok := r.Get("load"); !ok {
		t.Fatal("missing")
	}
}
