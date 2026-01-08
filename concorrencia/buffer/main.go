package main

import (
	"fmt"
	"time"
)


func addNums(canal chan int){
	time.Sleep(time.Second * 3)
	canal <- 10
	time.Sleep(time.Second * 3)
	canal <- 20
	canal <- 30
	canal <- 40
	canal <- 50
	canal <- 60
	fmt.Println("cheguei nesse ponto")
}

func main(){
	canal := make(chan int, 3)
	
	go addNums(canal)

	for i := range canal {
		fmt.Println(i)
	}
}