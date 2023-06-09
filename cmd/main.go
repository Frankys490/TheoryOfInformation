package main

import (
	GUI "TheoryOfInformation/pkg/GUI"
	Code "TheoryOfInformation/pkg/code"
	"fmt"
)

func main() {
	fmt.Println("Выберете лабораторную:")
	var numOfLab string
	fmt.Scan(&numOfLab)
	switch numOfLab {
	case "1":
		Code.Lab1()
	case "2":
		GUI.GuiLab2()
	case "4":
		GUI.GuiLab4()
	}
}
