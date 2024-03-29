package main

import (
	"fmt"
	"math"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 99}
	primeString := "is prime"
	notPrimeString := "is not prime"
	fmt.Println(isPrime(199))
	for _, num := range primes {
		if(isPrime(num)) {
			fmt.Println(num, primeString)
		} else {
			fmt.Println(num, notPrimeString)
		}
	}
	//binarySearchSortedArray(primes)
}

//function to determine if a number is prime
func isPrime(n int) bool {
	if(n <= 2) {
		return true
	}
	root := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 2; i <= root; i++ {
		if(n % i == 0) {
			return false
		}
	}
	return true
}

func binarySearchSortedArray(nums[] int){
	const start = 0
	end := len(nums)-1


	for start <= end {
		//middle := math.Floor(float64((start+end)/2))


	}
}
