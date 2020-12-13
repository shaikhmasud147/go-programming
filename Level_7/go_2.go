package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main()  {
	//Pointer demo
	x := person{
		fname: "Masud",
		lname: "Shaikh",
	}
	fmt.Println(x)
	changeMe(&x)
	fmt.Println(x)
}

func changeMe(p *person)  {
	p.fname = "Rahul"
	p.lname = "Sawant"
}