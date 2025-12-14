package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	fmt.Println("First-class citizens in Go. Declared with func keyword, support parameters and return values. Can be assigned to variables, passed as arguments,\nreturned from other functions. Support multiple return values, named returns, and variadic parameters. Building blocks of modular code. in details.")

	//! simple function
	var message string = greeting("john cena")
	fmt.Println("\n",message)

	// ! Variadic Functions
	fmt.Println("\nVariadic Functions ")
	fmt.Println("->. Functions accepting variable number of arguments of same type. Syntax: func name(args ...Type). Arguments treated as slice inside function. Call with multiple args or slice with ... operator. Common in functions like fmt.Printf() and append().")
	var tags []string = []string{"info", "debug", "error"}
	logger(tags...)

	//! Multiple Return Values
	fmt.Println("\nMultiple Return Values")
	fmt.Println("->.Go functions can return multiple values, commonly used for returning result and error. Syntax: func name() (Type1, Type2). Caller receives all returned values or uses blank identifier _ to ignore unwanted values. Idiomatic for error handling pattern.")
	

	//! VMultiple Returns
	fmt.Println("\nMultiple Return Values")
	fmt.Println("->.Go functions can return multiple values, commonly used for returning result and error. Syntax: func name() (Type1, Type2). Caller receives all returned values or uses blank identifier _ to ignore unwanted values. Idiomatic for error handling pattern.")
	data := make(map[string]interface{})
	data["name"] = "john cena"
	data["age"] = 48
	data["world_champ"] = 17
	response, err := commonResponse(data); if err != nil { fmt.Println("some thing went wrong.", err); return }
	fmt.Println("api call is success.", response)


	// !Anonymous Functions
	fmt.Println("\nAnonymous Functions")
	fmt.Println("->. Functions declared without names, also called function literals or lambdas. Can be assigned to variables, passed as arguments, or executed immediately. Useful for short operations, callbacks, goroutines, and closures. Access enclosing scope variables. Common in event handlers and functional patterns.")
	// !Anonymous Functions
	func(){
		fmt.Println("\nInside anonymous function")
	}()

	// !Closures
	fmt.Println("\nClosures")
	fmt.Println("->. Functions that capture and remember state from their enclosing scope. Can access variables from outer function even after outer function has finished execution. Common in higher-order functions, callbacks, and maintaining state in concurrent environments.")
	adder := adder()
	fmt.Println("adder(1):", adder(1))
	fmt.Println("adder(2):", adder(2))
	fmt.Println("adder(3):", adder(3))

	// !Named return values func
	fmt.Println("\nNamed return values func")
	fmt.Println("->. Functions that return named values. Syntax: func name() (Type1, Type2). Caller receives all returned values or uses blank identifier _ to ignore unwanted values. Idiomatic for error handling pattern.")
	success, message := namedReturnValues()
	fmt.Println("success:", success)
	fmt.Println("message:", message)

	// !Call by Value
	fmt.Println("\nCall by Value")
	fmt.Println("->. Go creates copies of values when passing to functions, not references to originals. Applies to all types including structs and arrays. Provides safety but can be expensive for large data. Use pointers, slices, maps for references. Critical for performance optimization.")
	const num int = 10
	fmt.Println("actual num:", num)
	callByValue(num)
	fmt.Println("outside callByValue num:", num)
}


// a function that return string //! simple function
func greeting(name string) string {
	return "Hello " + name + " good day to you!"
}


// ! Variadic Functions //Logger with Tags (Real-World Usage)
func logger(tags ...string){
	fmt.Println("Tags:")
	for _, tag := range tags {
		fmt.Println("-", tag)
	}
}

// this func return apis common response or error
// return type struct
type CommonResponse struct {
	success bool
	message string
	data interface{}
}
func commonResponse(data interface{}) (CommonResponse, error) {
	result, _ := json.Marshal(data)
	return CommonResponse {
		success: true,
		message: "api call is success.",
		data: string(result),
	}, nil
}

// !Closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// !Named return values func
func namedReturnValues() (success bool, message string) {
	return true, "api call is success."
}


// !Call by Value
func callByValue(num int) {
	num = num + 1;
	fmt.Println("inside callByValue num:", num)
}