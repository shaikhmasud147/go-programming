package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(a[:5])
	fmt.Println(a[5:10])
	fmt.Println(a[2:7])
	fmt.Println(a[1:6])

}
