package main

import "fmt"

func main() {
	a1 := []string{"pune", "satara", "sangli", "mumbai", "nasik"}
	a2 := []string{"kholapur", "nanded", "nagpur", "aurangabad", "nagar"}
	fmt.Println(a1)
	fmt.Println(a2)
	a3 := [][]string{a1, a2}
	for _, l := range a3 {
		fmt.Printf("\n")
		for j, val := range l {
			fmt.Printf("\n Index: %v\t value: %v", j, val)
		}
	}

}
