package main

import "fmt"

func main()  {
	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("Program complited")
}

// send data
func foo(c chan<- int)  {
	c <- 43
}

// recieve data
func bar(c <-chan int)  {
	fmt.Println(<-c)	
}