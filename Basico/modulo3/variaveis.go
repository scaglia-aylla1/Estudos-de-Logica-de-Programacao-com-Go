package main

import "fmt"

func main(){
	//Variaveis => var é usado mais a contexto de escopo global
	var nome = "Ana"
	fmt.Println(nome)
    
	//reatribuição
	nome = "Luna"
	fmt.Println(nome)

	var idade int
	idade = 20
	fmt.Println(idade)

	// declarando duas variaveis
	var b, c int = 1, 4
	fmt.Println(b)
	fmt.Println(c)

	var logado = true
	fmt.Println(logado)

	//Declarando com o operador curto => usado em escopo de bloco, declara e já atribui o valor
	f := "Luana"
	fmt.Println(f)

	numero := 55
	fmt.Println(numero)

	//Constantes => são variaveis imutavéis , ou seja não pode ser reatuibuido o valor
	const idadeAna = 30
	fmt.Println(idadeAna)

}