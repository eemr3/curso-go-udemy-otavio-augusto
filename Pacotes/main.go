package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do pacote main!")
	auxiliar.Escrevendo()
	error := checkmail.ValidateFormat("archimonder@gmail.com")
	fmt.Println(error)
}
