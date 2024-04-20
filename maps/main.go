package main

import "fmt"

func main() {
	fmt.Println("learing maps today")

	// we can intialise the map in two diffrent ways

	// way 1
	studentsGrade := make(map[string]int)

	// way2
	newStudentsAge := map[string]int{
		"Khan":  22,
		"Amir":  21,
		"Yo Yo": 55,
	}

	// adding the key value pair
	studentsGrade["Ankit"] = 88
	studentsGrade["Shubh"] = 78
	studentsGrade["Tel"] = 68
	studentsGrade["Paji"] = 58
	studentsGrade["Honey"] = 48

	// updating the grade
	studentsGrade["Shubh"] = 10000

	// deleting the key value pair
	delete(studentsGrade, "Tel")

	// checking the exsistence key
	// when we check the existing key in an array , then it return the value of it and if it is presnet then value will be the actual number other wise 0
	// now if it is present then it will return , the boolean value "T" , other wise false

	Garde, Exsits := studentsGrade["Paji"]

	Garde1, Exsits1 := studentsGrade["Tel"]

	fmt.Println("grade of paji", Garde)
	fmt.Println("is paji presnet", Exsits)

	fmt.Println("grade of paji", Garde1)
	fmt.Println("is paji presnet", Exsits1)
	fmt.Println("the students and grade is ", studentsGrade)
	fmt.Println("the students and grade is ", newStudentsAge)

}
