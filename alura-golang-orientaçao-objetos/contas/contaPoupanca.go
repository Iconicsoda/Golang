package contas

import (
	"Golang/alura-golang-orientaçao-objetos/clientes"
	"fmt"
)

//ContaPoupanca struct
type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

//Sacar faz o saque da contaCorrente desejada
func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	fmt.Println(&c)
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

//Depositar faz o deposito
func (c *ContaPoupanca) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito <= 0 {
		return "Deposito não foi possível", ValorDoDeposito
	} else {
		c.saldo += ValorDoDeposito
		return "Deposito feito com sucesso", ValorDoDeposito
	}
}

//ObterSaldo obtem o saldo
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
