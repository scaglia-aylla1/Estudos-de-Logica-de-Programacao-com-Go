package main

import "fmt"

// Funções em Go 
// Função começando com letra minúscula: É uma função Privada
//Só pode ser utilizada no próprio pacote


//Função começando com letra maiúscula: É uma função Pública
//Pode ser utilizada em outros pacotes
func main() {
	somaDosValores := soma(42, 13)
	fmt.Println(somaDosValores)

	nome, sobrenome := printaNomeCompleto("Aylla", "Scaglia")
	fmt.Println(nome)
	fmt.Println(sobrenome)
}


func soma(x int, y int) int{
	return x + y
}

func printaNomeCompleto(nome, sobrenome string) (string, string){
	return nome, sobrenome
}

