package main

import "fmt"

// input parameter data type would be float64, it means decimal place can be large ex- 10.3659635861552
//retun output parametr would be also be the same , it will also be in float

//(float64, error), first part would be the data , and 2nd part would be error , it is returning
// the retruning in the error can be string also
func divide(a, b float64) (float64, error) {
	// now checking the value of b

	if b == 0 {
		return 0,
			fmt.Errorf("2nd number should not be zero")

	}
	return a / b, nil // 1st part returning the data  && 2nd part returning the error

}

func main() {

	fmt.Println("Learning error handeling today ")

	result, err := divide(10, 0)

	// by comparing the error with nil , we can actually print the error
	if err != nil {
		fmt.Println(err)
	}
	// result will return the result
	// _ is used to ignore the error

	fmt.Println("final result would be ", result)
}
