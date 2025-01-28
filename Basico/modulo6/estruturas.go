package main

import "fmt"

func main(){
	//Array é uma coleção de valores,
    //arrays tem o tamanho definido em sua criação e eles são de tamanho fixo

	// pode ser declarado com var seguido do [tamanho]tipo
    var animes [3]string
    animes[0] = "Gantz"
    animes[1] = "Berserk"
    animes[2] = "Attack on Titan"
    fmt.Println(animes)

    // também pode utilizar o operador ':=' e passar os valores entre as chaves
    personagens := [3]string{"Kurono", "Guts", "Eren"}
    fmt.Println(personagens)

    // ou substituindo o tamanho por '...'
    personagens = [...]string{"Kei Kishimoto", "Casca", "Mikasa"}
    fmt.Println(personagens)


    //Slices são muito similares aos Arrays, porém são dinâmicos
    //a forma mais simples de criar um Slice é utilizando a função make:

    // make recebe o tipo []string e a capacidade inicial do slice
    // cria um slice com tamanho 3
     animes2 := make([]string, 3)
     animes2[0] = "Naruto"
     animes2[1] = "Kakashi"
     animes2[2] = "Sailor Moon"
     fmt.Println(animes2)
     fmt.Println(len(animes2), cap(animes2))
 
     // a função append redimensiona a capacidade do slice
     animes2 = append(animes2, "Claymore")
     fmt.Println(animes2)
     fmt.Println(len(animes2), cap(animes2))

    //Maps: Heterogêneos, ou seja, pode misturar tipos
	// estrutura de chave-valor => [key] = value
	//chave pode ter um tipo , eo valor pode ter outro tipo
    // map[string]int => como declarar um map
	//Exemplo => { "Ana": 28, "Pedro": 38}
    //map[string]string
	//Exemplo => { "Ana": "Silva", "Pedro": "Santana"}

     // são iniciados com make, é opcional informar a capacidade inicial
    // e assim é definido map[tipo da chave]tipo do valor
    dicionario := make(map[string]string)
    dicionario["hi"] = "oi"
    dicionario["bye"] = "tchau"
    dicionario["hey"] = "opa"
    fmt.Println(dicionario)

    // também podem ser iniciados com valores literais
    personagemAnime := map[string]string{
        "Eren":   "Tailor moon",
        "Guts":   "Naruto",
        "Kurono": "Claymore",
    }
    fmt.Println(personagemAnime)

}