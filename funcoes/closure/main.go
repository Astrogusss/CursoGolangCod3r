package main

import "fmt"

func closure () func(){
	x := 20
	funcao := func() {
		fmt.Println(x)
	}

	return funcao
}

func main(){
	x := 10
	fmt.Printf("Print da primeira vez: ")
	fmt.Println(x)

	//a funcao tem "lembranca" da onde ele foi chamado, mesmo com o mesmo nome da variável
	imprimeX := closure()
	fmt.Printf("Print da segunda vez: ")
	imprimeX()

}