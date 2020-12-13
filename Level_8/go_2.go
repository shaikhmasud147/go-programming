package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string `json: Fname`
	Lname string `json: Lname`
	Age int `json: Age`
}

func main()  {
	p := `[{"Fname":"Masud","Lname":"Shaikh","Age":29},{"Fname":"Asad","Lname":"Shaikh","Age":27}]`
	bs := []byte(p)

	fmt.Println("\n%T: ", p)
	fmt.Println("\n%T: ", bs)

	var peoples []person
	
	err := json.Unmarshal(bs, &peoples)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(peoples)

	for i, v := range peoples {
		fmt.Println("Index of row", i)
		fmt. Println("Person details: ",v.Fname, v.Lname, v.Age)
	}
}