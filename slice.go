package main

import "fmt"

func main() {

	var x []int

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 4)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
