// Package totient provides Euler's totient function.
package totient

// Phi TODO: implement Euler's product formula to compute phi.
func Phi(n int) int {
	res := 1
	for i := 2; i < n; i++ {
		if gcd(i, n) == 1 {
			res++
		}
	}
	return res
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
