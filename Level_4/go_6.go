package main

import "fmt"

func main() {
	a := make([]string, 5)
	a = []string{"pune", "satara", "sangli", "mumbai", "nasik"}
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	fmt.Println(len(a))
	fmt.Println(cap(a))
}
