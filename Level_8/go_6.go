package main

import (
	"fmt"
	"sort"
)

func main()  {
	randomNum := []int{2,4,6,1,4,8,7,9}

	sort.Ints(randomNum)
		
	fmt.Println(randomNum)

	city := []string{"Kholapur", "Pune", "Nasik", "Mumbai", "Satara"}

	sort.Strings(city)

	fmt.Println(city)
}

