package main

import "fmt"

func main() {
	a := 34
	fmt.Printf("%d\t%b\t%T", a, a, a)

	b := a << 1
	fmt.Printf("%d\t%b\t%T", b, b, b)

}
