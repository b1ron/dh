// Package mathutil provides some math utilities and helper functions.
package mathutil

import (
	"math/rand"
)

const maxUint32 = uint32(2<<31 - 1) // 4294967295

const bitsize = 2048
const numInts = bitsize / 32

// bigInt is a big integer type.
type bigInt struct {
	data [numInts]uint32
}

func (b *bigInt) String() string {
	// TODO use radix method
	return ""
}

// Add adds two *bigInt a and c together using bitwise addition.
// It stores the result in a larger type to prevent overflow.
func (b *bigInt) Add(a, c *bigInt) [numInts]uint64 {
	var res = [numInts]uint64{}
	for i, x := range a.data {
		y := c.data[i]
		res[i] = add(x, y)
	}
	return res
}

// Pow computes x^y using modular exponentiation.
func Pow(x, y, mod int) int {
	res := 1
	x = x % mod
	for y > 0 {
		if y%2 == 1 {
			res = (res * x) % mod
		}
		x = (x * x) % mod
		y /= 2
	}
	return res
}

// Pow computes x^y using modular exponentiation on a *bigInt.
func (b *bigInt) Pow(y *bigInt, mod uint32) uint32 {
	var res uint32 = 1
	for i, x := range b.data {
		x = x % mod
		for y.data[i] > 0 {
			if y.data[i]%2 == 1 {
				res = (res * x) % mod
			}
			x = (x * x) % mod
			y.data[i] /= 2
		}
	}
	return res
}

func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}

// ProbabilisticPrimeTest uses the Miller–Rabin algorithm to test if n is prime.
// https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
// It runs in polynomial time.
func ProbabilisticPrimeTest(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	a := 0
	d := n - 1
	s := 0
	x := 0

	for d%2 == 0 {
		d /= 2
		s++
	}

	for range 10 {
		a = rand.Intn(n-3) + 2
		// compute a^d mod n
		x = Pow(a, d, n)
		// handle later
		if x == 1 || x == n-1 {
			continue
		}
		for range s {
			x = Pow(x, 2, n)
			if x == n-1 {
				break
			}
		}
		if x != 1 && x != n-1 {
			return false
		}
	}

	// if none of the bases a detect compositeness, we assume n is probably prime.
	return true
}

func GeneratePrime(bitsize int) *bigInt { return nil }

func add(a, b uint32) uint64 {
	x := uint64(a)
	y := uint64(b)
	for y > 0 {
		carry := x & y
		x = x ^ y
		y = carry << 1
	}
	return x
}

func _() {
	// LSB refers to the least significant 32 bits, an element of data.
	// if LSB is 1, then it is prime
	// if LSB is 0, then it is composite
}
