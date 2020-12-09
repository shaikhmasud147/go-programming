package main

import "fmt"

func main() {
	intVal := []int{1, 2, 3, 4, 5, 6}
	result1 := fooSum(intVal...)
	fmt.Println(result1)
	result2 := barSum(intVal)
	fmt.Println(result2)

}

func fooSum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func barSum(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
