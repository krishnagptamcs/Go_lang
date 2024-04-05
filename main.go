// in a directory only one package is jst like main
// if you want to create another package then create a sub directory

package main

import (
	"fmt"
	"myLearining/utils" // using the package from utils
)

func main() {
	fmt.Println("lo bhai ho gaya print...")

	utils.PrintMessage("this is printing from utils package")

	//Learnig variable

	var name string = "krishna"
	fmt.Println(name)

	var lname = 2808
	fmt.Println(lname)

	const age int = 25
	// age = 58 // cannot assign any value to age , this is decleare then decleare
	fmt.Println(age)

	// without using any var/const and , without declaring its type
	//another practise

	fullName := "krishna gupta" // most of the varaible decleartion in go is done like this
	fmt.Println(fullName)

	var Private = "this data is private"
	var public = "this data is public"

	//Now if any variable decleare with capital letter , ex Privat , capital "P" then this varible can be export and it is used in another file , accessbale anywhere in the file
	// If any varriable is declearr in normal case , then it is used within the file , acceessable in the same file
	fmt.Println(Private)
	fmt.Println(public)

	// func PublicFunction() {
	// 	fmt.Println("this can be use inside the same file")
	// }

	// func privateFunction() {
	// 	fmt.Println("this can be use inside the same file")
	// }

}
