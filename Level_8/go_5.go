	package main

	import (
		"encoding/json"
		"fmt"
		"os"
	)

	type user struct {
		Name string
		Address string
		Mobile int
	}

	func main()  {
		u1 := user{
			Name: "Masud Shaikh",
			Address: "Pune",
			Mobile: 1234567890,
		}

		u2 := user{
			Name: "Satish Rathod",
			Address: "Mumbai",
			Mobile: 1234567890,
		}

		u3 := user{
			Name: "karan zohar",
			Address: "Mumbai",
			Mobile: 1234567890,
		}

		u4 := user{
			Name: "Sajit Nadiawala",
			Address: "Mumbai",
			Mobile: 1234567890,
		}

		users := []user{u1, u2, u3, u4}

		fmt.Println(users)

		err := json.NewEncoder(os.Stdout).Encode(users)

		if err != nil {	
			fmt.Println(err)
		}
	}