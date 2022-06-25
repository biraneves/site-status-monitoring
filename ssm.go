package main

import (
	"bufio" // Leitura do conteúdo de um arquivo
	"fmt"
	"io"
	"net/http" // Acesso à web via protocolo HTTP
	"os"       // Funcionalidades de interação com o sistema operacional
	"strings"
	"time" // Funções envolvendo tempo

	sv "github.com/biraneves/screen-visual"
)

const version = "0.1.0"
const scans = 5
const delay = 10


func main() {

	for {

		showMenu()
		option := getOption()

		switch option {

		case 0:
			os.Exit(0)

		case 1:
			startMonitoring()

		case 2:
			fmt.Println("\nExibindo logs")

		default:
			fmt.Println("\nOpção inválida")

		}

	}

}

func showMenu() {

	title := fmt.Sprint("Site Status Monitor - v. ", version)
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

	var sites []string
	
	fmt.Println("\nMonitorando...")
	fmt.Println("")

	sites = readSitesFromFile()

	for i := 0; i < scans; i++ {

		for _, site := range sites {

			testSite(site)
	
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

}

func testSite(site string) {

	resp, err := http.Get(site)
	statusCode := resp.StatusCode

	if err != nil {

		fmt.Println("Erro:", err)

	}

	fmt.Print("Testando site:", site)

	if statusCode == 200 {

		fmt.Println("\tOK")

	} else {

		fmt.Println("\tERRO\tstatus code:", statusCode)

	}

}

func readSitesFromFile() []string {

	var sites []string
	
	file, err := os.Open("sites.lst")

	if err != nil {

		fmt.Println("Erro:", err)

	}

	reader := bufio.NewReader(file)

	for {

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {

			break

		}

	}

	file.Close()

	return sites

}