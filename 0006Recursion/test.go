package main

import "fmt"

func main() {

	//test n!
	var temp = Factorial(20)
	fmt.Println("Result of n! = ", temp)

	//test fibonacci
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

func Factorial(n uint64) (result uint64) {
	//Function to calculate n!
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
