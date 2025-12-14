package main

import "fmt"

func main(){
	fmt.Println("Pointers in Go")
	fmt.Println("->. Variables storing memory addresses of other variables. Enable efficient memory usage and allow functions to modify values. Declared with [ *Type ] , address obtained with [ & ]. No pointer arithmetic for safety. Essential for performance and building data structures.")
	pointerExample()
}


func pointerExample(){
	var name = "jhon cena"
	var pointer_value = &name;
	fmt.Println("\nname has a value:", name)
	fmt.Println("\npointer_value has a value of memory address of name:", pointer_value)
	fmt.Println("\nwe did'nt assine pointer any direct value we just point what value should pointer_value will have and we extract from name which is :", *pointer_value)
}