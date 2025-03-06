package mathutil

import "testing"

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
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{997, true},
		{999, false},
		{1000, false},
		{1001, false},
	}
	for _, test := range tests {
		if got := ProbalisticPrimeTest(test.n); got != test.want {
			t.Errorf("probalisticPrimeTest(%d) = %t, want %t", test.n, got, test.want)
		}
	}
}
