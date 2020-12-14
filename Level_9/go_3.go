package main

import "fmt"

type person struct {
	fname string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySometing(h human)  {
	h.speak()
}

func main()  {
	p := person{
		fname: "Masud",
	}

	saySometing(&p)

	p.speak()
}