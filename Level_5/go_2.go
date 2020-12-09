package main

import "fmt"

func main() {
	t1 := struct {
		first    string
		friends  map[string]int
		favDrink []string
	}{
		first: "Masud",
		friends: map[string]int{
			"Asad": 555,
			"Q":    777,
			"M":    888,
		},
		favDrink: []string{
			"Appi",
			"Sprite",
		},
	}

	fmt.Println(t1)

}
