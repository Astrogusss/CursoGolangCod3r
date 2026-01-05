package main

import (
	"fmt"
	"time"
)

func multiplica345 (canal chan int, num int){
	time.Sleep(2 * time.Second)
	canal <- num * 3
	
	time.Sleep(2 * time.Second)
	canal <- num * 4

	time.Sleep(2 * time.Second)
	canal <- num * 5
}

func main(){
	var canal chan int
	canal = make(chan int)

	go multiplica345(canal, 10)

	a := <-canal
	fmt.Println(a)

	b := <-canal
	fmt.Println(b)

}