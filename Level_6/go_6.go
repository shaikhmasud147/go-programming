package main

import "fmt"

func main() {
	func() {
		x := 5
		y := 6
		z := x + y

		fmt.Println(z)
	}()
}
