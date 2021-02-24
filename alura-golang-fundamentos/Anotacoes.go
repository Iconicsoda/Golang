//arrays tem tamanhos fixos, sendo assim voce nao costuma usar arrays em go
//voce acaba utilizando uma coisa chamada slices que rodam emcima dos arrays

// como voce declara um array
// var sites [4]string

// FOR NORMAL
// for i := 0; i < len(sites); i++ {
// 	fmt.Println(sites[i])
// }

// //Slices
// func exibeNomes() {
// 	// voce tambem consegue gerar apenar um só, como no javascript
// 	// nomes := []string{}
// 	// Quando é necessário colocar mais elementos do que sua capacidade atual o slice dobra a capacidade
// 	nomes := []string{"Virto", "Nathan", "Dedeco", "Gustavo", "Diego"}
// 	nomes = append(nomes, "Endiabrado")
// 	fmt.Println(nomes)

// 	//para descobrir o length vc utiliza len(variavel)
// 	fmt.Println(len(nomes))

// 	//para descobrir a capacidade é cap(variavel)
// 	fmt.Println(cap(nomes))

// }

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
