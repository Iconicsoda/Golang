package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	//Golang nao tem o While, da para rodar um for indefinidamente
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)
	// fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
	}
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

// Igual no elixir vc consegue fazer que o go evite receber certos itens vindos de uma funcao
// apenas utilizando o _
// _, idade := devolveNomeEIdade()
// fmt.Println(idade)
// func devolveNomeEIdade() (string, int) {
// 	nome := "vitor"
// 	idade := 27
// 	return nome, idade
// }
