package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	leSitesDoArquivo()

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
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("")
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}
	// sites[0] = "https://random-status-code.herokuapp.com/"
	// sites[1] = "https://www.alura.com.br"
	// sites[2] = "https://www.caelum.com.br"

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		// tipo um map da vida LOL (posicao, quem ta na posiçao)
		for i, site := range sites {
			fmt.Println("Testando site:", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
	}

}

func leSitesDoArquivo() []string {
	var sites []string

	// arquivo, err := ioutil.ReadFile("sites.txt")

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//Retorna um leitor
	leitor := bufio.NewReader(arquivo)

	for {
		// '' significa bytes
		linha, err := leitor.ReadString('\n')

		// Corta o final dos espaços
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Ocorreu um erro na leitura:", err)
		}

		//Esse string ele é tipo um ToString
		// fmt.Println(string(linha))
	}

	//Sempre uma boa pratica FECHAR O TXT ou o arquivo
	arquivo.Close()

	return sites
}
