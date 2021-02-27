//Mesmo não provendo nenhum valor,
//o Go garante inicializar todas as variáveis, conforme a imagem abaixo:
//
// Boll -> false
// int -> 0
// float -> 0
// string -> ""
// struct {}
//
//

// Função variadics, elas conseguem por varios ints em sequencia
// func Somando(numeros ...int) int {
//     resultadoDaSoma := 0
//     for _, numero := range numeros {
//         resultadoDaSoma += numero
//     }
//     return resultadoDaSoma
// }

// func main() {
//     fmt.Println(Somando(1))
//     fmt.Println(Somando(1,1))
//     fmt.Println(Somando(1,1,1))
//     fmt.Println(Somando(1,1,2,4))
// }

// antigo code
// contaDaSilvia := new(c.ContaCorrente)
// contaDaSilvia.Titular = "Silvia"
// contaDaSilvia.Saldo = 500

// contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 1000}

// statusDepositar := contaDaSilvia.Transferir(100, &contaDoGustavo)

// fmt.Println(statusDepositar)

// fmt.Println(contaDoGustavo)

// contaDaSilvia.Sacar(-100)

// status, valor := contaDaSilvia.Depositar(1000)
// fmt.Println(status, valor)
// fmt.Println(contaDaSilvia.Saldo)