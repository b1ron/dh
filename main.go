package main

import (
	"fmt"

	"github.com/b1ron/dh/mathutil"
	"github.com/b1ron/dh/totient"
)

// https://math.stackexchange.com/questions/124408/finding-a-primitive-root-of-a-prime-number

// A number g is called a primitive root modulo n if:
// g^k mod n produces all the numbers relatively prime to n before repeating, for k = 1, 2, 3, ..., φ(n), where φ is Euler's totient function.
func main() {
	// primitive root modulo 7 is 3
	seen := make(map[int]bool)
	primitiveRoot := 3
	for i := 1; i < totient.Phi(7); i++ {
		g := mathutil.Pow(primitiveRoot, i, 7)
		if seen[g] {
			break
		}
		seen[g] = true
	}
	// 3, 2, 6, 4, 5, 3
	fmt.Println(primitiveRoot)
}
