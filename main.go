package main

import "fmt"

func GetMember() {
	fmt.Println("Please wait....")
}

func IsInSystem(username string) bool {
	return true
}

func GetUserDetail(username string) (int, string) {
	return 201, "manager"
}

func getDeparture(username string, departure *string) {
	if username != "" {
		*departure = "home"
	}
}

func CheckLogin(username string, password string) {
	if IsInSystem(username) {
		fmt.Println("found user in system.")
		status, deatail := GetUserDetail(username)
		fmt.Printf("status %d detail %s\n", status, deatail)
		departure := ""
		getDeparture(username, &departure)
		fmt.Println(departure)
	}
}

func main() {
	GetMember()
	CheckLogin("coregate", "1234")
}
