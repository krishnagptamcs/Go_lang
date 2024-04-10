package main

import "fmt"

func main() {
	fmt.Println("Learnig Arrays today ")

	var name [5]string

	name[0] = "mr"
	name[1] = "krishna"
	name[3] = "gupta"

	// another way of declearing the array
	var numbers = [5]int{1, 2, 3, 4, 5}

	//In Go, when you declare an array or a slice, the elements are initialized to their zerovalues. The zero value is the default value assigned to variables of a certain type when no explicit value is provided.

	//For numeric types (int, float, etc.), the zero value is 0. For strings, it is an empty string(""). For boolean types, it is false, and for pointers or complex types, it is nil.

	fmt.Println("name array contains", name)
	fmt.Println("number array contains", numbers)

	// length of array , len(numbers)
	fmt.Println("length of number array is ", len(numbers))
	fmt.Println("length of number array is ", len(name))
}
