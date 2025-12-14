package main

import (
	"encoding/json"
	"fmt"
)

type Vertex struct {
	Salary int
	Name string
}


type User struct {
	UserName string
	Email string
	IsLogin bool
	IsActive bool
}


func main() {
	fmt.Println(Vertex{Salary: 6500, Name: "Zaheer shaikh"})
	user_1 := User{
		UserName: "zaheer shaikh",
		Email: "zaheer.zonesso@gmail.com",
		IsLogin: true,
		IsActive: false,
	}
	fmt.Print("==================>")
	fmt.Println(user_1)
	fmt.Print("==================>")
	// map is diffrent map["key"]
	// but in struct we use [stuck].[key]
	fmt.Println(user_1.Email)
	fmt.Print("==================>")
	user_2 := User{
		Email: "zaheer.shaikh@zonesso.com",
	}
	fmt.Println(user_2)
	fmt.Print("==================>")
	result, err := json.Marshal(user_1); if err != nil {fmt.Println(err)};
	fmt.Println(string(result))

}
