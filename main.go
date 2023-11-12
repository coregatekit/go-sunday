package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Address  string
	PostCode string
}

type UserProfile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int
	Height    float32

	Address Address

	Bill struct {
		BillAddress string
	}
}

func (u UserProfile) ToFullName() string {
	// return u.Firstname + " " + u.Lastname
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

func main() {
	fmt.Println("Map Struct Inteface")

	// Map
	user := map[string]string{} // create empty map object with key: string and value type string
	user["username"] = "dekjong"
	user["password"] = "super_dekjong"
	fmt.Println(user)
	fmt.Println(user["username"])
	fmt.Println(user["password"])

	// Struct
	userProfile := UserProfile{
		Firstname: "Joe",
		Lastname:  "Kim",
		Age:       26,
		Height:    176.9,
		Address: Address{
			Address:  "123 Tower",
			PostCode: "10900",
		},
		Bill: struct{ BillAddress string }{
			BillAddress: "123 Tower A",
		},
	}

	fmt.Println(userProfile)
	fmt.Println(userProfile.Firstname)
	fmt.Println(userProfile.Lastname)
	fmt.Println(userProfile.ToFullName())

	byteTxtJson, err := json.MarshalIndent(userProfile, "", " ")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteTxtJson))

	dataJson := `
	{
		"firstname": "Wendy",
		"lastname": "Shon",
		"Age": 29,
		"Height": 158,
		"Address": {
		 "Address": "123 Tower",
		 "PostCode": "10900"
		},
		"Bill": {
		 "BillAddress": "123 Tower A"
		}
	}
	`
	var wendyProfile UserProfile
	err = json.Unmarshal([]byte(dataJson), &wendyProfile)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wendyProfile)
}
