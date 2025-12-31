package main

import (
	"fmt"
)

func analisarMedia(media func(num1, num2 float32) float32, num1, num2 float32) {
	if resultado := media(num1, num2); resultado >= 6 {
		fmt.Println("Você passou")
	} else {
		fmt.Println("Reprovado")
	}
}

func main(){

	num1, num2 := 5, 5.5
	media := func(num1, num2 float32) float32{
		return (num1 + num2) / 2
	}

	analisarMedia(media, float32(num1), float32(num2))
}
