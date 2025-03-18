// Package mathutil provides some math utilities and helper functions.
package mathutil

import (
	"fmt"
	"math/rand"
)

// XXX use 64-bit chunks instead of 32-bit - fewer operations per (*,+), and more efficient storage.
const maxUint32 = uint32(2<<31 - 1) // 4294967295
const maxUint64 = uint64(2<<63 - 1) // 18446744073709551615

const maxBitsize = 2048
const numInts = maxBitsize / 32

// bigInt is a dynamically sized integer type.
type bigInt struct {
	data []uint32 // dynamically sized
}

func newBigInt(bitsize int) (*bigInt, error) {
	if bitsize <= 0 || bitsize > maxBitsize {
		return nil, fmt.Errorf("bitsize is out of range: %d", bitsize)
	}
	numInts := (bitsize + 31) / 32 // + 31 is needed to round up and not truncate
	return &bigInt{data: make([]uint32, numInts)}, nil
}

func (b *bigInt) String() string {
	return ""
}

// Add adds two *bigInt a and c together using bitwise addition.
// It stores the result in a larger type to prevent overflow.
func (b *bigInt) Add(a, c *bigInt) [numInts]uint64 {
	return [numInts]uint64{}
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

// This is the meat of how we compute an arbitrarily long prime number, given a bitsize.
// It uses an entropy source to generate a random number for each chunk.
// It uses the Miller-Rabin primality test to check if the number is prime.
func GeneratePrime(bitsize int) *bigInt { return nil }

func _(a, b uint32) uint64 {
	x := uint64(a)
	y := uint64(b)
	for y > 0 {
		carry := x & y
		x = x ^ y
		y = carry << 1
	}
	return x
}
