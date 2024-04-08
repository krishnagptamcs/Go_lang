package main

import "fmt"

func simpleFunction() {
	fmt.Println("another function call inside main function")
}

// declaring input parameter type and return type
func addTwoNumbers(a, b int) int {
	return a + b
}

// declearing input parameter and output parameter type with result
func addAnotherTwoNumber(a, b int) (result int) {
	result = a + b
	return
}
func main() {
	fmt.Println("learning function today")

	simpleFunction()

	result := addTwoNumbers(5, 8)

	finalResult := addAnotherTwoNumber(10, 20)

	fmt.Println("the result of two number is ", result)
	fmt.Println("the result of two number is ", finalResult)
}
