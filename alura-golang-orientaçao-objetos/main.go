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

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	fmt.Println(&c)
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(ValorDoDeposito float64) string {
	if ValorDoDeposito <= 0 {
		return "Deposito não foi possível"
	} else {
		c.saldo += ValorDoDeposito
		return "Deposito feito com sucesso"
	}
}

func main() {
	contaDaSilvia := new(ContaCorrente)
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	contaDaSilvia.Sacar(-100)

	fmt.Println(contaDaSilvia.saldo)
}
