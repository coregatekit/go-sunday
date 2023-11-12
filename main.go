package main

import "fmt"

type SaveLogInterface interface {
	SaveLog()
}

func SaveLog(st SaveLogInterface) {
	st.SaveLog()
}

type TransferBBL struct {
	Name string
}

func (tBBL *TransferBBL) SaveLog() {
	tBBL.Name = "test"
	fmt.Println("save to database", tBBL.Name)
}

type TransferKTB struct {
	Name string
}

func (tKTB TransferKTB) SaveLog() {
	fmt.Println("save to database", tKTB.Name)
}

func main() {
	fmt.Println("Map Struct Inteface")

	transA := TransferBBL{
		Name: "BBL",
	}
	transB := TransferKTB{
		Name: "KTB",
	}
	SaveLog(&transA)
	SaveLog(transB)

	fmt.Println(transA.Name)
}
