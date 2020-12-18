package main

import "fmt"

func main()  {
	c := make(chan int)

	go foo(c)

	for i := range c {
		fmt.Println(i)
	}

}

func foo(c chan<- int)  {
	for i:=0; i<=10; i++{
		c <- 42
	}

	close(c)
}