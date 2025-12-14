package main

import (
	"fmt"
	"sort"
)

func main() {
	array_of_int := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array_of_int)
	var arr_length int = len(array_of_int) // this is length
	fmt.Println(arr_length)
	var index_of_zero int = array_of_int[0]
	fmt.Println(index_of_zero)
	 
	fmt.Println("============>")
	//! Array Slicing in Golang
	sort.Ints(array_of_int[:])
	fmt.Println(array_of_int[0:])
	
	fmt.Println("============>")
	//! Array Sorting
	fmt.Println(array_of_int)
}