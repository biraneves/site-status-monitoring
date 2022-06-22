package main

import (
	"fmt"
	"os"		// Funcionalidades de interação com o sistema operacional
	"net/http"	// Acesso à web via protocolo HTTP
	sv "github.com/biraneves/screen-visual"
)

const VERSION = "0.1.0"

func main() {

	showMenu()
	option := getOption()

	switch option {

	case 0:
		os.Exit(0)

	case 1:
		startMonitoring()

	case 2:
		fmt.Println("Exibindo logs")

	default:
		fmt.Println("Opção inválida")
		os.Exit(-1)

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

func startMonitoring() {

	fmt.Println("\nMonitorando...")

	site := "http://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)
	statusCode := resp.StatusCode

	if statusCode == 200 {

		fmt.Println("Site:", site, "carregado com sucesso!")

	} else {

		fmt.Println("Site:", site, "apresentou problema - status code:", statusCode)

	}

}