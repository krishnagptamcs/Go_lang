package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What is your name dear?")

	// var name string //Declares a variable named name of type string.

	// fmt.Scan(&name) //Reads input from the standard input and stores it in the memory location referenced by the address of the name variable.

	// // through this we can read before upto white space

	// fmt.Println("Hello Mr.", name)

	// Now to take input and scan whole word including  white space

	reader := bufio.NewReader(os.Stdin) // bufio is a package which read upto whole line
	// os.stdin --> operating system ka standard input

	name, _ := reader.ReadString('\n') // read tb tk kro jab tk next lin e na mil jaaye

	//A buffered reader is a type of reader that reads data from an underlying source, such asa file or standard input (keyboard), and stores that data in a buffer. The buffer is atemporary storage area in memory. Buffered readers are commonly used to improve theefficiency of input operations.

	fmt.Println("Hello Mr.", name)

	//In Go, the & operator is used to obtain the memory address of a variable. It is called the "address of" operator.

	//In the provided code snippet, fmt.Scan(&name) is used to read input from the standard input (keyboard) and store it in the variable name. However, Scan function in Go requires the memory address of the variable where the input should be stored. Hence, &name is used to pass the memory address of the name variable to the Scan function.
}
