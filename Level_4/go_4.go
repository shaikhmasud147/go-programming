package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(a)

	//Append data
	a = append(a, 52)
	fmt.Println(a)

	a = append(a, 53, 54, 55)
	fmt.Println(a)

}
