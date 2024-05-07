package main

import "fmt"

// making the structure od the name  Person
// defing its name and age

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Number int
	Email  string
}

type Address struct {
	HouseNumber int
	Area        string
	Pincode     int
}

// creating a another structure , which follow the sturcture type
type Employee struct {
	Person_Details Person  // name  string
	Person_Contact Contact // gender string
	Person_Address Address // age int
}

func main() {

	var empDetails Employee

	empDetails.Person_Details = Person{
		FirstName: "new krishna",
		LastName:  "new gupta",
		Age:       25,
	}

	empDetails.Person_Contact.Number = 1234567890
	empDetails.Person_Contact.Email = "k@gmail.com"
	empDetails.Person_Address = Address{
		HouseNumber: 8,
		Area:        "xyz chandigarh india",
		Pincode:     12346,
	}

	fmt.Println("the employee details are ", empDetails)
	fmt.Println("the employee details are ", empDetails.Person_Details)
	fmt.Println("the employee details are ", empDetails.Person_Contact)
	fmt.Println("the employee details are ", empDetails.Person_Address)
	// taking a variable krishna
	// and assing to the structure person
	var krishna Person
	fmt.Println("learning data structure today ")

	// fmt.Println("Person Krishna :", krishna)

	// now assing the value in the structure
	//type 1

	krishna.FirstName = "krishna"
	krishna.LastName = "gupta"
	krishna.Age = 60
	fmt.Println("the Value of krishna after filling the data is ", krishna)

	// assing the value in type 2
	// here the structure is using of the person

	krishna2 := Person{
		FirstName: "Krishna1",
		LastName:  "Gupta1",
		Age:       30,
	}

	fmt.Println("the value of krishna 2", krishna2)
	// type 3

	krishna3 := new(Person)

	krishna3.FirstName = "Krishna 3"
	krishna3.LastName = "Gupta 3"
	krishna3.Age = 50

	fmt.Println("krihsna3 data is ", krishna3)

}
