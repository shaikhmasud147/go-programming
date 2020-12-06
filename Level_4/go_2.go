package main

import "fmt"

func main() {
	a := []int{4, 5, 2, 6, 1, 23, 45, 76}

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("\n%T", a)
}
