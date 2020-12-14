package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {
	
	pass := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pass)
	fmt.Println(string(bs))

	loginPass := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))

	if err != nil {
		fmt.Println("You can't login")
		return 
	}

	fmt.Println("You are logged in")
}