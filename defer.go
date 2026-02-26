package main

import "fmt"

func sum(a int, b int) (s int) {
	s = a + b
	return
}

func main() {

	r := sum(5, 10)

	fmt.Println(r)

}
