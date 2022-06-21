package main

import (
	"fmt"
	sv "github.com/biraneves/screen-visual"
	"reflect"
)

func main() {

	nome := "Bira"
	idade := 49
	versao := 1.1
	fmt.Println("Olá, sr.", nome, "!")
	fmt.Println("Sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	sv.Line("-", 64)
	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))

}
