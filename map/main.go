package main

import "fmt"

func main(){
	user := map[string] interface{} {
		"name": "zaheer shaikh",
		"age":30,
		"salary":6500,
	}

	fmt.Println("==================")
	fmt.Println("\n",user)
	fmt.Println("==================")
	fmt.Println("\n",user["salary"])
}
