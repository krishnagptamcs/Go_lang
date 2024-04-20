package main

import "fmt"

func main() {
	fmt.Println("learing for loop today")
	fmt.Println("in go lang there is no any other loop , only for loop is present ")

	for i := 0; i < 4; i++ {
		fmt.Println("maza aawata.....")
	}

	// running infite loop
	counter := 0

	for {
		fmt.Println("inifite loop trail")
		counter++

		if counter == 8 {
			break
		}
	}

	// range key word

	// range is used to find the value and index of an array

	numbers := []int{1, 2, 3, 4, 5}

	// index , and value of array numbers
	for index, value := range numbers {
		fmt.Printf("index :%d , value: %d\n", index, value)

	}

}
