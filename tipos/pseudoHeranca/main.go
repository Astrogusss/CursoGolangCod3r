package main

import "fmt"

type carro struct{
	nome string
	velocidade float32
}

type ferrari struct{
	carro
	turbo bool
}

func main(){
	//carro1 := ferrari{carro{"Spider", 100}, true}
	carro2 := ferrari{}

	carro2.nome = "Spider"
	carro2.velocidade = 100
	carro2.turbo = true

	fmt.Printf("O nome do carro é %s e que tem velocidade é %f com o seu turbo %v \n", carro2.nome, carro2.velocidade, carro2.turbo)
}