package main

import "fmt"

func modifyByPointers(num *int) {
	// getting the value of that address by *
	// changing the value by addinjg +20  in it
	*num = *num + 20
}

func main() {
	fmt.Println("learing pointer today")
	//Pointers--> they provide a way to work the memory directly , for effecient memory managment and sharing data b/w functions
	// Pointer --> sirf varaible ke memory ka address store karta hai , it does not mean with its variable value

	// var num int
	// num = 2

	// taking a pointer ptr , and the value which is stored in this address its data type is integer
	// var ptr *int

	// ptr == address of num variable , address represent by &
	// ptr = &num

	// method 2 of storing

	num := 2
	ptr := &num

	fmt.Println("the value of num :", num)
	fmt.Println("the value of pointer", ptr) //0xc00000a0b8 , it is address of the block where the num varaible is stored
	fmt.Println("data contains on pointer address is ", *ptr)

	value := 10

	ptrAdd := &value

	// function which modifys the value by pointer ,
	modifyByPointers(ptrAdd)

	fmt.Println("value of variable is", value)

}
