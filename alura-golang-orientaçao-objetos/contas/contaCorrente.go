package contas

import (
	"Golang/alura-golang-orientaçao-objetos/clientes"
	"fmt"
)

//ContaCorrente struct
type ContaCorrente struct {
	//se voce escreve com minusculo
	//signifca que esse valor é um valor privado, que nao pode ser acessado fora desse package
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	//se eu deixo com a letra minuscula, significa que é privado, sendo assim ele aparece undefined para as outras
	saldo float64
}

//Sacar faz o saque da contaCorrente desejada
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	fmt.Println(&c)
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

//Depositar faz o deposito
func (c *ContaCorrente) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito <= 0 {
		return "Deposito não foi possível", ValorDoDeposito
	} else {
		c.saldo += ValorDoDeposito
		return "Deposito feito com sucesso", ValorDoDeposito
	}
}

//Transferir transfere xD
func (c *ContaCorrente) Transferir(ValorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if ValorDaTransferencia < c.saldo && ValorDaTransferencia > 0 {
		c.saldo -= ValorDaTransferencia
		contaDestino.Depositar(ValorDaTransferencia)
		return false
	} else {
		return true
	}
}

//ObterSaldo -> se vc colocar letras minusculas na função
//ela vai ser uma funçao privada xD
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
