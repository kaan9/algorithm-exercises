package main

import "fmt"

func gcd(a uint, b uint) uint {
	for a > 0 {
		a, b = b, a
		a %= b
	}
	return b
}

func main() {
	fmt.Println(gcd(674,234))
}
