package libs

import "fmt"

func Greeting() {
	fmt.Println("Hi there!")
}

func Login(username string, password string) bool {
	if username == "admin" && password == "1234" {
		return true
	}
	return false
}
