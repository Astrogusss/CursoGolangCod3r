package main

import (
	"fmt"
	"time"
)

func NumerosPares (canal chan int){
	inicio := 2

	for i := 0; i < cap(canal) * 2 ; i++ {
		for aux := inicio; ; aux++{
			if aux % 2 == 0 {
				time.Sleep(time.Second * 2)
				canal <- aux
				inicio = aux + 1
				break
			}
		}
	}
	//o close vai servir para que o for que esta na main seja sinalizado que nao vai ser enviado mais nada para o determinado canal
	close(canal)
}


func main(){
	canal := make(chan int, 4)
	
	go NumerosPares(canal)

	for i := range canal {
		fmt.Println(i)
	}

}