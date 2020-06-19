package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func hasError() (int, bool) {
	return 0, false
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	_, error := hasError()
	fmt.Println(error)
}
