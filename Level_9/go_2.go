package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main()  {
	
	fmt.Println("Go CPU\t ", runtime.NumCPU())
	fmt.Println("Goroutine\t ", runtime.NumGoroutine())
	wg.Add(1)
	
	go foo()
	go bar()

	wg.Wait()
	fmt.Println("Go CPU\t ", runtime.NumCPU())
	fmt.Println("Goroutine\t ", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i<=10; i++ {
		fmt.Println(i)
	}

	wg.Done()
}

func bar() {
	for i := 0; i<=10; i++ {
		fmt.Println(i)
	}

	wg.Done()
}