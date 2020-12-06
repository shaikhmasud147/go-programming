package main

import "fmt"

func main() {

	//Key value data
	m := map[string][]string{
		`profile`: []string{`First Name`, `Last Name`, `City`, `Phone`},
		`company`: []string{`Masud`, `Shaikh`, `Pune`, "0123456789"},
	}

	for k, v := range m {
		fmt.Println(k, v)
		for k, v := range v {
			fmt.Println(k, v)
		}
	}
}
