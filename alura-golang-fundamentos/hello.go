package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	//No go quando vc nao atribui valor a uma var ele retorna 0 ou ""
	var nome string = "Virto"
	//var idade = 12
	versao := 1.1 //isso daqui tbm funciona pq ele pega o tipo e ainda coloca o tipo
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

//Golang só aceita se receber apenas operadores que mandam true ou false
// if comando == 1 {
// fmt.Println("Monitorando...")
// } else if comando == 2 {
// 	fmt.Println("Exibindo logs...")
// } else if comando == 0 {
// 	fmt.Println("Saindo do programa, FONTRAB")
// } else {
// 	fmt.Println("Não conheço esse comando")
// }
