package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak(s string) {
	fmt.Printf("%v %v: %v\n", p.first, p.last, s)
}

func main() {
	p1 := person{
		first: "Masud",
		last:  "Shaikh",
		age:   29,
	}

	p2 := person{
		first: "Sachin",
		last:  "Piolat",
		age:   27,
	}

	p1.speak("loud")
	p2.speak("slow")
}
