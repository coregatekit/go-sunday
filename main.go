package main

import "fmt"

func main() {
	// arr := []int{2, 4, 4, 5}
	arr := make([]int, 4) // create empty array of int with size 4
	arr[0] = 30
	arr[2] = 50
	fmt.Println(arr)

	txt := "today is sunday"
	// slice
	fmt.Println(txt[0:5])
	fmt.Println(arr[0:1])
	fmt.Println(len(txt))
}
