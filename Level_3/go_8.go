package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Should not be print")
	case true:
		fmt.Println("Should be print")
	}
}
