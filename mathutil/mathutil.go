// Package mathutil provides some math utilities and helper functions.
package mathutil

// helper to avoid annoying float64 conversion
func Pow(x, y int) int {
	res := 1
	for range y {
		res *= x
	}
	return res
}

func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}
