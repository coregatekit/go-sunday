package main

import "fmt"

func main() {
	// fmt.Println("enter your name: ")
	// var txtInput string
	// txtInput := ""
	// fmt.Scan(&txtInput)
	// fmt.Println("" + txtInput)

	// age := 2
	// fmt.Println(reflect.TypeOf((age)))

	// height := "100"

	// // var h int
	// // var err error
	// h, err := strconv.Atoi(height)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(h)

	// water := "12.40"
	// fWater, err := strconv.ParseFloat(water, 64)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(fWater)

	// txtShow := fmt.Sprintf("water = %8.2f, height = %d\n", fWater, h)
	// fmt.Print(txtShow)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// score := 80
	// switch score {
	// case 80:
	// 	{
	// 		fmt.Println("A")
	// 	}
	// case 70:
	// 	{
	// 		fmt.Println("B")
	// 	}
	// }

	score := 77
	if score >= 80 {
		fmt.Println("A")
	} else if score < 80 && score >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("unknow")
	}
}
