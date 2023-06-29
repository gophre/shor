package shor

import (
	"math"
	"math/rand"
)

// Shor's algorithm to factorize a number N
func Factorize(N int) (int, int) {
	if N%2 == 0 {
		return 2, N / 2
	}

	for {
		a := rand.Intn(N-2) + 2
		gcdVal := gcd(a, N)
		if gcdVal > 1 {
			return gcdVal, N / gcdVal
		}

		r := findPeriod(a, N)
		if r%2 != 0 {
			continue
		}

		x := int(math.Pow(float64(a), float64(r/2))) % N
		if x == (N - 1) {
			continue
		}

		p := gcd(x+1, N)
		q := gcd(x-1, N)
		if p*q == N {
			return p, q
		}
	}
}

// Extended Euclidean algorithm to find the greatest common divisor (gcd) of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Quantum period finding subroutine
func findPeriod(a, N int) int {
	initialState := rand.Intn(N-1) + 1
	x := 1

	// Apply modular exponentiation to the initial state
	for i := 0; i < initialState; i++ {
		x = (x * a) % N
	}

	return x
}
