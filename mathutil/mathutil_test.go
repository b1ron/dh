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
	// TODO:
	// test 2^2048 % (2^2048 - 1) = 1
}

func TestBigInt(t *testing.T) {
	// makes a bigInt of size 2^2048
	b, err := newBigInt(2048)
	assert.NoError(t, err)
	// logarithmic transformation from base 2 to base 10
	// log_{10}(2^2048) = log{10}(10^e)
	// 2048 * log_{10}(2) = e * log_{10}(10)
	// e = 2048 * log_{10}(2)
	// e = 2048 * log_{10}(2) = 2048 * 0.3010 = 616.509
	// so 2^2048 has 617 digits, and requires 64 32-bit integers
	assert.Equal(t, len(b.data), 64)

	split := func(s string, n int) []string {
		chunks := make([]string, 0, n)
		chunkSize := len(s) / n
		remainder := len(s) % n

		start := 0
		for i := range n {
			extra := 0
			if i < remainder {
				// need to distribute the extra character
				extra = 1
			}
			end := start + chunkSize + extra
			chunks = append(chunks, s[start:end])
			start = end
		}
		return chunks
	}

	// 2^2048 ~=
	s := "32317006071311007300714876688669951960444102669715484032130345427524655138867890893197201411522913463688717960921898019494119559150490921095088152386448283120630877367300996091750197750389652106796057638384067568276792218642619756161838094338476170470581645852036305042887575891541065808607552399123930385521914333389668342420684974786564569494856176035326322058077805659331026192708460314150258592864177116725943603718461857357598351152301645904403697613233287231227125684710820209725157101726931323469678542580656697935045997268352998638215525166389437335543602135433229604645318478604952148193555853611059596230656"
	chunks := split(s, 64)
	assert.Equal(t, len(chunks), 64)

	// WIP
	b.GeneratePrime()
	t.Log(b.data)
}
