package main

import "fmt"

func main() {
	// ! custom type conversion
	type MyTypeOfString string

	var name MyTypeOfString = "Zaheer shaikh"
	fmt.Println(name)
	fmt.Printf("type of name: %T", name)


	//! Function type
	var result GreetingT = GreetingT(greeting)
	var greeting_result string = result("\n Zaheer shaikh, Welcome to go lang")
	fmt.Println(greeting_result)

	var f GreetingT = GreetingT(greeting) 
	fmt.Println(f)

}


type GreetingT func(string) string

func greeting(value string) string {
	return "hello " + value
}