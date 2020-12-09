package main

import "fmt"

func main() {
	a := 20
	b := 30
	resultInt := foo(a, b)
	fmt.Println(resultInt)

	data := "foobar"
	resultString := bar(data)
	fmt.Println(resultString)
}

func foo(x int, y int) int {
	c := x + y
	return c
}

func bar(s string) string {
	stringData := s
	return stringData
}
