package main

import "fmt"

type masud int

var x masud

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println(x)
}
