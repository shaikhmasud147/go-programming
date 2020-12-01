package main

import "fmt"

type masud int

var x masud
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 20
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
}
