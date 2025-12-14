package main

import "fmt"

type Employee struct {
	Name string
	Company string
	Salary int
}

func main(){
	person := new(Employee)
	person.Name = "Zaheer"
	person.Company = "Zonesso"
	person.Salary = 6500

	fmt.Println(person)

	// ! make function

	fmt.Print("==================>")
	a := make([]int, 6)
	fmt.Println(a)
	fmt.Print("==================>")
	a = append(a, 1)
	fmt.Println(a)
}