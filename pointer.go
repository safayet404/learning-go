package main

import "fmt"

func print(numbers *[3]int) {
	fmt.Println(numbers)
}

func main() {

	// x := 10

	// p := &x

	// *p = 20

	// fmt.Println("Address", p)
	// fmt.Println("Value at address", *p)

	arr := [3]int{1, 2, 3}
	print(&arr)

}
