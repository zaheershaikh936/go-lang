package main

import "fmt"

func main() {
	x := 10;
	fmt.Println("Outer variable x: ", x)
	if true {
		x := 20;
		fmt.Println("Inner variable x: ", x)
	}
	fmt.Println("Outer variable x: ", x)
}