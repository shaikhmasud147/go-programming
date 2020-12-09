package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	fic1 := person{
		first: "Masud",
		last:  "Shaikh",
	}

	fic2 := person{
		first: "Karan",
		last:  "Johar",
	}

	m := map[string]person{
		fic1.first: fic1,
		fic2.first: fic2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
	}
}
