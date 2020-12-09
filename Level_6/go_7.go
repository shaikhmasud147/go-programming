package main

import "fmt"

func main() {
	f := func() {
		x := 5
		y := 6
		z := x + y

		fmt.Println(z)
	}

	f()

	fmt.Printf("%T", f)
}
