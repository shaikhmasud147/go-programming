package main

import (
	"runtime"
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i:=0; i<=gs; i++ {
		go func ()  {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End value: ", incrementer)
}