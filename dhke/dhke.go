// Package dhke provides a simple Diffie-Hellman key exchange implementation.
package dhke

import (
	"math"

	"github.com/b1ron/dh/mathutil"
)

// probalisticPrimeTest uses the Miller–Rabin algorithm to test if n is prime.
// https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
// It runs in polynomial time.
func probalisticPrimeTest(n int) bool {
	a := 0
	d := 3
	s := 0
	x := 0
	y := 0
	for range n - 1 {
		s = mathutil.Pow(2, s)
		if s*d == n-1 {
			a = s
			break
		}
	}

	// [2, min(n − 2, ⌊2(ln n)^2⌋)]
	bound := math.Min(float64(n-2), math.Floor(2*math.Pow(math.Log(float64(n)), 2)))
	for range int(bound) {
		// compute a^d mod n
		x = mathutil.Pow(a, d) % n
		// handle later
		if x == 1 || x == n-1 {
			break
		}
		for range s {
			y = mathutil.Pow(x, 2) % n
		}
		if y == 1 && x != 1 && x != n-1 {
			return false
		} else {
			x = y
			break
		}
	}

	// if none of the bases aa detect compositeness, we assume n is probably prime.
	return y == 1
}
