package main

import "fmt"

func main() {
	// s := []int{1, 2, 5}

	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	s := make([]int, 3, 5)

	s[0] = 10
	s[2] = 50

	fmt.Println(s)

}
