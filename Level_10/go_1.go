package main

import "fmt"

func main()  {
	c := make(chan int, 2)

	c <- 20
 	c <- 30
	
	fmt.Println(<-c)
	fmt.Println(<-c)
}