package main

import (
	"fmt"
	t "fundamentos/teste"
)

func main(){
	t.Escrever()
	t.Escrever2()
	
	var(
		peidao int8 = 10
		peidao2 int8 = 11
	)	
	fmt.Println("valor do peidao é", peidao, "e o outro valor é", peidao2)
}
