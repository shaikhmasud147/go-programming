package main

import (
	"fmt"
)

//global variables
var x int = 40
var y string = "James Bond"
var z bool = true

func main() {
	//print varibales with new line
	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z)
	fmt.Println(s)
}
