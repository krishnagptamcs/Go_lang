// learning slice

package main

import (
	"fmt"
)

func main() {
	// fmt.Println("learing slice today")
	// fmt.Println("slice in nothing a new thing it is same just like an array , the main diff is that , you do not have define its size while declearing so that adding/ deleting more data can be done ")

	// numbers := []int{1, 2, 3, 4, 5, 6}

	// numbers = append(numbers, 55, 60, 88, 96, 36, 15)

	// fmt.Println("the slice is ", numbers)                             // to print the slice
	// fmt.Printf("the slice number has data type is : %T\n  ", numbers) // to print its data type
	// fmt.Println("length of the slice is ", len(numbers))              // to print its length

	// creating slice by "make"

	age := make([]int, 3, 6)

	// here -->[]int  is an datatype of the slcie
	// here -->3 is an length of the slcie
	// here -->6 is an capity of the slcie  (max storage )
	// but since it is dynamic , so if the capacity is full and you filling more data to it , then its capity will increase by 2x

	age = append(age, 50, 60, 20)
	age = append(age, 50, 60, 20, 69)

	fmt.Println("Slice", age)
	fmt.Println("Length:", len(age))
	fmt.Println("Capacity:", cap(age))

}
