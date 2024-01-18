package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Hello, my name is Mogly.")
	fmt.Println("Welcome my first programe in Go.")

	// Variaveis em Go
	var nome string = "Mogly"
	var idade int = 25
	var versao float32 = 1.1
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	// Variaveis em Go - Inferencia de tipo
	var nome2 = "Mogly"
	var idade2 = 25
	var versao2 = 1.1
	fmt.Println("O tipo da variavel nome2 é", reflect.TypeOf(nome2))
	fmt.Println("O tipo da variavel idade2 é", reflect.TypeOf(idade2))
	fmt.Println("O tipo da variavel versao2 é", reflect.TypeOf(versao2))

	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O valor da variável comando é:", comando)
}
