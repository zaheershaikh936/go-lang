package main

import "fmt"

func main(){
	// var i , j int = 20, 2
	// k := 3
	// fmt.Println("this is i: ", i , "this is j: ", j, "this is k: ", k)
	// java_script, python, java, go_lang := true, false, false, true;
	// fmt.Printf("this is java_script: %t /n this is python: %t /n this is java: %t /n this is go_lang: %t", java_script, python, java, go_lang)
	
	// var interger int;
	// var string_value string;
	// var isBoolean bool;
	// fmt.Printf("%v\n", interger)
	// fmt.Printf("%v\n", string_value)
	// fmt.Printf("%v\n", isBoolean)

	// ! const
	
	//? Multiple constants:
	const (
		User = "user"
		ShowRoom = "showroom"
		Admin = "admin"
	)
	const user string = ShowRoom;
	fmt.Println("this is user type: ", user)
}