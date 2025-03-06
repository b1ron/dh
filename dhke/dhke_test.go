package dhke

import (
	"testing"

	"github.com/b1ron/dh/mathutil"
)

func TestProbalisticPrimeTest(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{997, true},
	}
	for _, test := range tests {
		if got := mathutil.ProbalisticPrimeTest(test.n); got != test.want {
			t.Errorf("probalisticPrimeTest(%d) = %t, want %t", test.n, got, test.want)
		}
	}
}
