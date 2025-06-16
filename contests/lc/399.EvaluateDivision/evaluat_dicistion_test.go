package _399_evaulate_division

import (
	"reflect"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		expexted  []float64
	}{
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			expexted:  []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
	}
	for _, tt := range tests {
		result := calcEquation(tt.equations, tt.values, tt.queries)
		if !reflect.DeepEqual(result, tt.expexted) {
			t.Errorf("Got: %v expected: %v", result, tt.expexted)
		}
	}

}
