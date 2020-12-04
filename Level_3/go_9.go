package main

import "fmt"

const (
	a int = 20
	b int = 30
	c int = 0
)

func main() {
	i := "Multiplication"

	switch i {
	case "Addition":
		c := a + b
		fmt.Println(c)
	case "Multiplication":
		c := a * b
		fmt.Println(c)
	case "Dividetion":
		c := a / b
		fmt.Println(c)
	default:
		fmt.Println(c)
	}
}
