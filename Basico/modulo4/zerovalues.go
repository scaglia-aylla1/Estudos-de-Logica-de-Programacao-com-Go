package main

import "fmt"


func main(){

	// Conceito Zero values => quando a variavel não é iniciada com um valor,
	// o Go tem um valor Default
	var i int      // => valor 0
	var f float64  // => valor 0.0
	var b bool     // => valor bool sempre começa com falso
	var s string   // => valor vazio ""

	fmt.Printf("Inteiro: %v\n", i)
	fmt.Printf("Float: %v\n", f)
	fmt.Printf("Bool: %v\n", b)
	fmt.Printf("String: %q\n", s)




}