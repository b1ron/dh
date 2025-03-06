package main

import (
	"fmt"

	mathutil "github.com/b1ron/dh/math"
	"github.com/b1ron/dh/totient"
)

// https://math.stackexchange.com/questions/124408/finding-a-primitive-root-of-a-prime-number

// A number g is called a primitive root modulo n if:
// g^k mod n produces all the numbers relatively prime to n before repeating, for k = 1, 2, 3, ..., φ(n), where φ is Euler's totient function.
func main() {
	// primitive root modulo 7 is 3
	m := map[int]struct{}{}
	primitiveRoot := 3
	for i := 1; i < totient.Phi(7); i++ {
		g := mathutil.Pow(primitiveRoot, i, 7)
		if _, ok := m[g]; ok {
			break
		}
		m[g] = struct{}{}
	}
	// 3, 4, 5, 6, 2, 1
	fmt.Println(primitiveRoot)
}
