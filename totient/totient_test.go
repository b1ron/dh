package totient

import "testing"

func TestPhi(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 4},
		{6, 2},
	}
	for _, test := range tests {
		if got := Phi(test.n); got != test.want {
			t.Errorf("Phi(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
