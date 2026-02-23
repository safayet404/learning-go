package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	var user1 User

	user1 = User{
		Name: "Safayet",
		Age:  26,
	}

	fmt.Println("name : ", user1.Name)
	fmt.Println("age : ", user1.Age)

	user2 := User{
		Name: "Roki",
		Age:  33,
	}

	fmt.Println("name : ", user2.Name)
	fmt.Println("age : ", user2.Age)

}
