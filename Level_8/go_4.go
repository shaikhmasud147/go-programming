package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Address string
	Mobile int
}

func main()  {

	u := `[{"Name":"Masud Shaikh","Address":"Pune","Mobile":1234567890},{"Name":"Satish Rathod","Address":"Mumbai","Mobile":1234567890},{"Name":"karan zohar","Address":"Mumbai","Mobile":1234567890},{"Name":"Sajit Nadiawala","Address":"Mumbai","Mobile":1234567890}]`
	bs := []byte(u)

	var users []user

	err := json.Unmarshal(bs, &users)

	if err != nil {	
		fmt.Println(err)
	}

	fmt.Println(users)

	for i, v := range users {
		fmt.Println(i)
		fmt.Println(v)
	}
}