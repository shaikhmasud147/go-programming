package main

import "fmt"

func main() {
	a := [5]int{3, 4, 5, 6, 1}

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", a)
}
