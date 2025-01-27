package main

import(
	"fmt"
)

func main(){
	// bool (true/false)
	fmt.Printf("Type: %T - Value: %v\n", true, true)

	// string - sequência de caracteres
	fmt.Printf("Type: %T - Value: %v\n", "aylla", "aylla")

	// int - numero inteiro
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)

	// float - numero decimal(float64/float32)
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)
	fmt.Printf("Type: %T - Value: %v\n", 0.1, 0.1)
}