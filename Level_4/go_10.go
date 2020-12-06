package main

import "fmt"

func main() {

	//Key value data
	m := map[string][]string{
		`profile`: []string{`First Name`, `Last Name`, `City`, `Phone`},
		`company`: []string{`Masud`, `Shaikh`, `Pune`, "0123456789"},
	}

	m["test"] = []string{`test1`, `test2`, `test3`, `test4`}

	for k, v := range m {
		fmt.Println(k, v)
		for k, v := range v {
			fmt.Println(k, v)
		}
	}

	delete(m, "company")
	fmt.Printf("\n")

	for k, v := range m {
		fmt.Println(k, v)
		for k, v := range v {
			fmt.Println(k, v)
		}
	}
}
