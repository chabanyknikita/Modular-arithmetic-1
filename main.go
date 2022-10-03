package main

import (
	"math"
	"math/rand"
	"time"
)

// task 1
func mod(mod int) int {
	return mod
}

func euler(n int) int {
	f := n
	if n%2 == 0 {
		for n%2 == 0 {
			n = n / 2
		}
		f = f / 2
	}
	i := 3
	for i*i <= n {
		if n%i == 0 {
			for n%i == 0 {
				n = n / i
			}
			f = f / i
			f = f * (i - 1)
		}
		i = i + 2
	}
	if n > 1 {
		f = f / n
		f = f * (n - 1)
	}
	return f
}

func gcd(a, m int) int {
	if a < m {
		a, m = m, a
	}
	for m != 0 {
		a, m = m, a%m
	}
	return a
}

// task 2
func task2(a, m int) int {
	return a % m
}

// task 3
func task3(a, b, m int) int {
	return int(math.Pow(float64(a), float64(b))) % m
}

// task 4
func task4(a, b, m int) int {
	step := int(math.Pow(float64(a), float64(euler(m)-gcd(a, m))))
	x := (b * step) % m
	return x
}

// task 5
func task5(a, b int) int {
	var sl []int
	for i := a; i <= b; i++ {

		isPrime := true

		for j := 2; j < i; j++ {

			if i%j == 0 {

				isPrime = false
			}
		}

		if isPrime == true {
			sl = append(sl, i)
		}
	}
	rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(len(sl))
	return sl[randIdx]
}

func main() {

}
