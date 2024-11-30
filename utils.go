package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func StartingProject() {
	myFigure := figure.NewFigure("File Manager", "", true)
	myFigure.Print()

	fmt.Println("===================================")
	fmt.Println("Bem-vindo ao File Manager!")
	fmt.Println("Inicializando o sistema...")
	fmt.Println("===================================")
}
