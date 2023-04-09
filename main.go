// Write a program that takes a list of integers as input and returns a new list containing only the prime numbers from the original list
package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func filterPrimes(nums []int) []int {
	primes := []int{}
	for _, num := range nums {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	primes := filterPrimes(nums)
	fmt.Println(primes)
}
