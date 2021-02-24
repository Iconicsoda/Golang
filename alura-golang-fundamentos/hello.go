package main

import (
	"fmt"
)

func main() {
	//No go quando vc nao atribui valor a uma var ele retorna 0 ou ""
	var nome string = "Virto"
	//var idade = 12
	versao := 1.1 //isso daqui tbm funciona pq ele pega o tipo e ainda coloca o tipo
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)

	//Golang só aceita se receber apenas operadores que mandam true ou false
	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa, FONTRAB")
	} else {
		fmt.Println("Não conheço esse comando")
	}
}
