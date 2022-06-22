package main

import (
	"fmt"
	sv "github.com/biraneves/screen-visual"
)

const VERSION = "0.1.0"

func main() {

	var option  int

	show_menu()
	fmt.Scan(&option)

}

func show_menu() {

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