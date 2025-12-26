package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GerarNumerosAleatórios () int {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := randGenerator.Intn(10 - 3 + 1) + 3
	fmt.Println(num)
	return num
}

func main(){
	//fazendo if com a inicializacao de uma variável, ficando valida somente dentro do escopo
	//muito comum para tratar erros, fazendo com as variáveis erros nao fiquem pinduradods fora do programa principal

	if numeroSortido := GerarNumerosAleatórios(); numeroSortido >= 7 {
		fmt.Println("esse número é maior/igual que 7")
	} else {
		fmt.Println("esse numero é menor que 7")
	}

}
