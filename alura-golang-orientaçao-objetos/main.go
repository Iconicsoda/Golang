package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoVirto := ContaCorrente{"virto", 123, 123456, 1235.35}
	contaDoVirto2 := ContaCorrente{titular: "virto", saldo: 1235.35}

	fmt.Println(contaDoVirto2)
	fmt.Println(contaDoVirto)

}
