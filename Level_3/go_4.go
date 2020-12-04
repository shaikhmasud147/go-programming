package main

import "fmt"

func main() {
	i := 1991
	for {
		if i > 2021 {
			break
		}
		fmt.Println(i)
		i++
	}
}
