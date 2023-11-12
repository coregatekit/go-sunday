package main

import (
	"fmt"
	"time"
)

func GetMember() {
	fmt.Println("Please wait....")

	time.Sleep(3 * time.Second)
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

func CheckServerResponse() {
	fmt.Println("check server time")
	time.Sleep(3 * time.Second)
	panic("server error")
}

func LogEnd() {
	now := time.Now()
	fmt.Println("complete program ")
	fmt.Println(now)
}

func main() {
	// GetMember()
	// CheckLogin("coregate", "1234")

	defer LogEnd()
	GetMember()
	CheckLogin("coregate", "123")
	CheckServerResponse()
}
