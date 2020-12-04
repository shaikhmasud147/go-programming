package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Number %v is even\n", i)
		} else {
			fmt.Printf("Number %v is odd\n", i)
		}
	}

}
