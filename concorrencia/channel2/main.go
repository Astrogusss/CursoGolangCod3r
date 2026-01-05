package main

import (
	"fmt"
	"time"
)

func SairMultiplicando(num int, canal chan int){
	for i := 0; i < 10; i++ {
		canal <- i * num
		time.Sleep(time.Second)
	}
}

func main(){
	canal := make(chan int, 1)

	go SairMultiplicando(3, canal)

	time.Sleep(time.Second * 10)
	fmt.Println(<-canal)
	

}