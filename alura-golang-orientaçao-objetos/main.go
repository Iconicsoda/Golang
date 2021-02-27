package main

import (
	// "Golang/alura-golang-orientaçao-objetos/clientes"
	c "Golang/alura-golang-orientaçao-objetos/contas"
	"fmt"
)

func main() {
	contaExemplo := c.ContaCorrente{}
	contaExemplo.Depositar(100.56)

	fmt.Println(contaExemplo.ObterSaldo())

	// contaDoBruno := c.ContaCorrente{
	// 	Titular: clientes.Titular{
	// 		Nome:      "Bruno",
	// 		CPF:       "111.111.111-11",
	// 		Profissao: "Programador",
	// 	},
	// 	NumeroAgencia: 1231,
	// 	NumeroConta:   11231231231,
	// 	Saldo:         1000,
	// }

}
