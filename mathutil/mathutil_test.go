package mathutil

import "testing"

func TestProbabilisticPrimeTest(t *testing.T) {
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
		if got := ProbabilisticPrimeTest(test.n); got != test.want {
			t.Errorf("ProbabilisticPrimeTest(%d) = %t, want %t", test.n, got, test.want)
		}
	}
}

func TestBigIntPow(t *testing.T) {
	maxInt := bigInt{}
	maxInt.data = [numInts]uint32{}

	// makes a bigInt of size 2^2048
	for i := range numInts {
		maxInt.data[i] = maxUint32
	}

	// test 2^2048 % (2^2048 - 1) = 1
	y := bigInt{}
	y.data = [numInts]uint32{}
	for i := range numInts {
		y.data[i] = maxUint32
	}

	got := maxInt.Pow(&y, maxUint32-1)
	if got != 1 {
		t.Errorf("bigInt.Pow(%d, %d) = %d, want %d", maxInt, maxUint32-1, got, 1)
	}
}

func TestAdd(t *testing.T) {
	b := bigInt{}

	x := bigInt{}
	x.data = [numInts]uint32{}
	for i := range x.data {
		x.data[i] = maxUint32
	}

	y := bigInt{}
	y.data = [numInts]uint32{}
	for i := range y.data {
		y.data[i] = maxUint32
	}

	want := [numInts]uint64{
		8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590,
		8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590,
		8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590,
		8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590,
		8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590,
		8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590,
		8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590,
		8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590, 8589934590,
	}
	if got := b.Add(&x, &y); got != want {
		t.Errorf("bigInt.Add(%d, %d) = %d, want %d", x, y, got, want)
	}
}
