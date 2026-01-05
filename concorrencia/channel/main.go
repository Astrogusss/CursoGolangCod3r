package main

import "fmt"

func main(){
	ch := make(chan int)
	
	//o channel nada mais é que uma porta de comunicacao
	//aqui estamos enviando i numero 1 para o canal de comunicacao (escrita)
	
	//Escrita
	ch <- 1
	
	//leitura
	i := <-ch

	//canal é um ponto de sincronismo entre goroutines, quando criamos as goroutines, quando criamos os canais, fazemos um ponto de intersessão entre esses goroutines concorrentes
	fmt.Println(i)
}