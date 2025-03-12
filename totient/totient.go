// Package totient provides Euler's totient function.
package totient

import "github.com/b1ron/dh/mathutil"

// https://en.wikipedia.org/wiki/Euler%27s_totient_function
// Phi TODO: implement Euler's product formula to compute phi.
func Phi(n int) int {
	res := 1
	for i := 2; i < n; i++ {
		if mathutil.GCD(i, n) == 1 {
			res++
		}
	}
	return res
}
