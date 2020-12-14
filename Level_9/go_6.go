package main

import (
	"runtime"
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	var incrementer int32
	gs := 100
	wg.Add(gs)

	var a sync.Atomic

	for i:=0; i<=gs; i++ {
		go func ()  {
			a.AddInt64(&incrementer, 1)
			fmt.Println(a.LoadInt64(&incrementer))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End value: ", incrementer)
}