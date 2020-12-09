package main

import "fmt"

func main() {

	defer foo()

	result2 := bar()
	fmt.Println(result2)

}

func foo() {
	fmt.Println("foo")
}

func bar() string {
	return "bar"
}
