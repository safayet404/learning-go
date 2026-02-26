package main

import "fmt"

func calculate() (result int) {

	fmt.Println("first", result)
	defer func() {
		result = result + 10

		fmt.Println("defer", result)
	}()

	result = 5

	fmt.Println("second result", result)

	return
}

func calc() int {

	result := 0

	fmt.Println("first", result)
	defer func() {
		result = result + 10

		fmt.Println("defer", result)
	}()

	result = 5

	fmt.Println("second result", result)

	return result

}
func main() {

	a := calculate()
	b := calc()

	fmt.Println(a)
	fmt.Println(b)

}
