// Package dhke provides a simple Diffie-Hellman key exchange implementation.
package dhke

import "math/rand"

// probalisticPrimeTest uses the Miller–Rabin algorithm to test if n is prime.
// https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
// It runs in polynomial time.
func probalisticPrimeTest(n int) bool {
	a := 2
	for ; a > 2; a = rand.Int() {
		if a < n-2 {
			break
		}
	}
	d := n - 1
	return false
}

func primeFactors(n int) []int {
	primes := []int{}
	divisors := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	for _, d := range divisors {
		for n%d == 0 {
			n /= d
		}
		primes = append(primes, d)
		if n == 1 {
			break
		}
	}
	return primes
}

// let s > 0 and d odd > 0 such that n − 1 = 2sd  # by factoring out powers of 2 from n − 1
// for all a in the range [2, min(n − 2, ⌊2(ln n)2⌋)]:
//     x ← ad mod n
//     repeat s times:
//         y ← x2 mod n
//         if y = 1 and x ≠ 1 and x ≠ n − 1 then  # nontrivial square root of 1 modulo n
//             return “composite”
//         x ← y
//     if y ≠ 1 then
//         return “composite”
// return “prime”
