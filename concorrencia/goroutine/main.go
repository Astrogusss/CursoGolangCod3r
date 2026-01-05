package main

import (
	"fmt"
	"time"
)

func falarMerdarPorAI(contagem int8, nome string){
	for i := 0; i < int(contagem); i++ {
			fmt.Printf("iteração merda %d quem disse foi %s \n", int(i + 1), nome)
			time.Sleep(time.Second)
	}
}

func main(){
	go falarMerdarPorAI(10, "Gustavo")
	go falarMerdarPorAI(7, "Cinara")
	time.Sleep(time.Second * 10)
}