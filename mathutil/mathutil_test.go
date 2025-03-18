package mathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	// maxInt := bigInt{}
	// maxInt.data = [numInts]uint32{}

	// // makes a bigInt of size 2^2048
	// for i := range numInts {
	// 	maxInt.data[i] = maxUint32
	// }

	// // test 2^2048 % (2^2048 - 1) = 1
	// y := bigInt{}
	// y.data = [numInts]uint32{}
	// for i := range numInts {
	// 	y.data[i] = maxUint32
	// }

	// got := maxInt.Pow(&y, maxUint32-1)
	// if got != 1 {
	// 	t.Errorf("bigInt.Pow(%d, %d) = %d, want %d", maxInt, maxUint32-1, got, 1)
	// }
}

func TestBigInt(t *testing.T) {
	_, err := newBigInt(-1)
	assert.Error(t, err)

	b, err := newBigInt(1024)
	assert.NoError(t, err)
	assert.Equal(t, len(b.data), 32)

	b, err = newBigInt(2048)
	assert.NoError(t, err)
	assert.Equal(t, len(b.data), 64)

	// 32317006071311007300714876688669951960444102669715484032130345427524655138867890893197201411522913463688717960921898019494119559150490921095088152386448283120630877367300996091750197750389652106796057638384067568276792218642619756161838094338476170470581645852036305042887575891541065808607552399123930385521914333389668342420684974786564569494856176035326322058077805659331026192708460314150258592864177116725943603718461857357598351152301645904403697613233287231227125684710820209725157101726931323469678542580656697935045997268352998638215525166389437335543602135433229604645318478604952148193555853611059596230656

}
