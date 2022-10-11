package main

import (
	"errors"
	"fmt"
)

func main() {
	// int => int8 int16 int32 int64 ou int
	var interiro8 int8 = 123
	fmt.Println(interiro8)
	var interiro16 int16 = 10000
	fmt.Println(interiro16)
	var interiro32 int32 = 1000000000
	fmt.Println(interiro32)
	var interiro64 int64 = 1000000000000000000
	fmt.Println(interiro64)
	numero := 1203545624
	fmt.Println(numero)

	// real (float) => float32 float64
	var real32 float32 = 100000.45
	fmt.Println(real32)
	var real64 float64 = 100000.125
	fmt.Println(real64)
	real := 1582.75
	fmt.Println(real)

	// string
	var str = "Texto"
	fmt.Println(str)
	str2 := "Texto 2"
	fmt.Println(str2)

	// char não existe dentro do Go ele referencia ao valor na tabela ASCII
	char := 'b'
	fmt.Println(char)

	// valores iniciais (zero)
	var str3 string
	var n int
	var r float64
	var verdOuFalse bool
	fmt.Println(str3, n, r, verdOuFalse)

	// tipo Error
	var erro error
	fmt.Println(erro)
	erro = errors.New("Erro inesperado!")
	fmt.Println(erro)

}
