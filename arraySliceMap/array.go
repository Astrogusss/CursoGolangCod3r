package main

import "fmt"

func ArraySlice(){
	var notas [3]int
	var media float64

	notas[0], notas[1], notas[2] = 12, 11, 10

	for chave, valor := range(notas){ 
		media += float64(valor) / 3
		fmt.Println(chave)
	}

	//aqui diz que o tamanho do vetor será contado a partir da quantidade de posicoes dos vetores colocados inicialmente, nao da pra colcoar depois
	// notasArray := [...]int{10, 11, 12}	

	//slice

	//slice nao é array, é uma fatia de um array interno
	notasArray := [3]int{7, 8, 9}

	//ele somente é um ponteiro
	notasFatiaArray := notasArray[0:2]

	fmt.Println("isso é o array principal", notasArray)
	fmt.Println("isso é a fatia do array principal", notasFatiaArray, "sua capacidade é de ", cap(notasFatiaArray), "e esta preenchido com as posicoes", len(notasFatiaArray))

	// notasSlice := []int{10, 11, 12}

	// fmt.Println("aqui temos ", cap(notasSlice), "e de tamanho temos ", len(notasSlice))

	// notasSlice = append(notasSlice, 13)

	// fmt.Println("aqui temos ", cap(notasSlice), "e de tamanho temos ", len(notasSlice))

}

func Maps(){

}

func main(){
	ArraySlice()
	Maps()

}