package main

import (
	"fmt"
	"time"
)

func Ping1S(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("ping : ", i)
		time.Sleep(1 * time.Second)
	}
	ch <- 10
}

func SendNoti5s(ch chan string) {
	fmt.Println("send noti...")
	time.Sleep(5 * time.Second)
	fmt.Println("Send noti...")
	ch <- "success"
}

func main() {
	fmt.Println("Go Routine & Channel")

	chPing := make(chan int)
	chNoti := make(chan string)

	go Ping1S(chPing)
	go SendNoti5s(chNoti)

	// time.Sleep(12 * time.Second)
	pingVal, notiMsg := <-chPing, <-chNoti
	fmt.Println(pingVal, notiMsg)
	fmt.Println("completed")
}
