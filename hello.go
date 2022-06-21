package main

import (
	"fmt"
	screenvisual "github.com/biraneves/screen-visual"
)

func main() {
	message := "Hello, World!"
	fmt.Println(message)
	screenvisual.Line("=", len(message))
}
