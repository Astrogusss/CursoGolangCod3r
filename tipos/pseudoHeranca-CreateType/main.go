package main

import (
	"fmt"
)

type carro struct{
	nome string
	velocidade float32
}

type ferrari struct{
	carro
	turbo bool
}

//criando nosso próprio tipo
type nota float32

//funcao para verificar se esta dentro do intervalo
func (n nota) ValidationBetween(inicio, fim int) bool{
	//tem que tipar para comprar, nao basta somente colocar float32, tem que ser do mesmo tipo, no caso, tem que ser nota
	if n >= nota(inicio) && n < nota(fim) {
		return true
	}
	return false
}

func (n nota) Conceito() string{
	if n.ValidationBetween(9, 10){
		return "A"
	} else if n.ValidationBetween(8, 9){
		return "B"
	} else if n.ValidationBetween(7, 8){
		return "C"
	} else {
		return "D"
	}
}

func main(){
	//carro1 := ferrari{carro{"Spider", 100}, true}
	carro2 := ferrari{}

	carro2.nome = "Spider"
	carro2.velocidade = 100
	carro2.turbo = true

	fmt.Printf("O nome do carro é %s e que tem velocidade é %f com o seu turbo %v \n", carro2.nome, carro2.velocidade, carro2.turbo)


	var prova nota = 9
	if validacao := prova.ValidationBetween(0, 10); validacao{
		fmt.Println("Sua nota esta coerente")
	} else {
		fmt.Println("Sua nota não esta coerente")
	}

	fmt.Println("Seu conceito é", prova.Conceito())

}