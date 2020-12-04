package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// module print 1 to 100
		fmt.Printf("When %v divided by 4, the reminder (aka module) is %v\n:", i, i%4)
	}
}
