package main

import (
	"fmt"
	"time"
)

func FalandoCanal(nome string) <-chan string{
	c := make(chan string)

	go func(){
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando o numero %d", nome, i)
		}
	}()

	return c
}

func juntar(canal1, canal2 <-chan string) <-chan string{
	resultado := make(chan string)


	go func(){
		for {
			select{
			case s := <-canal1: // caso o canal1 entre primeiro, ele ja vai colocar dentro do resutlado
				resultado <- s
			case s := <-canal2: // caso seja o canal2 o mais rapido, ele ja entra dentro do resultado
				resultado <- s
			}
		}
	}()

	return resultado
}


func main(){
	canal := juntar(FalandoCanal("Gustavo Senador"), FalandoCanal("Cinara Seandor"))

	fmt.Println(<-canal)
	fmt.Println(<-canal)
	fmt.Println(<-canal)
	fmt.Println(<-canal)

}