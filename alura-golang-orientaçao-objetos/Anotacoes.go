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