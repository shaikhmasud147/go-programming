package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	// Capital first later for JSON data
	Fname string
	Lname string
	Age	int
}

func main()  {
	// Json Marshal example
	p1 := person{
		Fname: "Masud",
		Lname: "Shaikh",
		Age: 29,
	}

	p2 := person{
		Fname: "Asad",
		Lname: "Shaikh",
		Age: 27,
	}

	peoples := []person{p1, p2}

	fmt.Println(peoples)

	bs, err := json.Marshal(peoples)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}