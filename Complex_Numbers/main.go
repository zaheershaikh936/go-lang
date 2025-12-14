package main

import "fmt"

func main() {
	//! constructoe init
	var c1 complex128 = complex(1, 2)
	fmt.Println(c1)

	//! complez number init sysntax
	var c2 complex128 = 10 + 11i 
	fmt.Println(c2)

	//! string_literals
	string_literals := "Go is expressive, concise, clean, and efficient. \n Its concurrency mechanisms make it easy to write programs \n that get the most out of multi-core and networked machines, \n while its novel type system enables flexible and modular \n program construction. Go compiles quickly to machine code \n yet has the convenience of garbage collection and the power \n of run-time reflection. It's a fast, statically typed, \n compiled language that feels like a dynamically typed, \n interpreted language."
	fmt.Println(string_literals)
}