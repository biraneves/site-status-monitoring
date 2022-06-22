package main

import (
	"fmt"
	sv "github.com/biraneves/screen-visual"
)

const VERSION = "0.1.0"

func main() {

	showMenu()
	option := getOption()

	switch option {

	case 0:
		fmt.Println("Saindo do programa")

	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Exibindo logs")

	default:
		fmt.Println("Opção inválida")

	}

}

func showMenu() {

	title := fmt.Sprint("Site Status Monitor - v. ", VERSION)
	fmt.Println("")
	fmt.Println(title)
	sv.Line("=", len(title))
	fmt.Println("[1] Iniciar monitoramento")
	fmt.Println("[2] Exibir logs")
	fmt.Println("[0] Sair")
	sv.Line("-", len(title))
	fmt.Print("Sua opção: ")

}

func getOption() int {

	var option int
	fmt.Scan(&option)

	return option

}