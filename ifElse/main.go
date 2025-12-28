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

func Conceito(nota float64) string{
	switch{
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 7 && nota < 8:
		return "C"
	case nota >= 6 && nota < 7:
		return "D"
	default:
		return "E"
	}
}

func main(){
	//fazendo if com a inicializacao de uma variável, ficando valida somente dentro do escopo
	//muito comum para tratar erros, fazendo com as variáveis erros nao fiquem pinduradods fora do programa principal

	if numeroSortido := GerarNumerosAleatórios(); numeroSortido >= 7 {
		fmt.Println("esse número é maior/igual que 7")
	} else {
		fmt.Println("esse numero é menor que 7")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	switch GerarNumerosAleatórios(){
	case 10, 9, 8, 7:
		fmt.Println("ofuhsduifhuisd")
	case 6:
		fallthrough
	case 5, 4, 3, 2, 1:
		fmt.Println("os outros")
	} 

	t :=  time.Now()
	fmt.Println(t)
	
	//quase como se fosse um if, mas so que esta passando pro switch um valor true como comeco, depois ele sai procurando por ai
	switch{
	case t.Hour() < 12:
		fmt.Println("bom dia")
	case t.Hour() < 18:
		fmt.Println("boa tarde")
	default:
		fmt.Println("boa noite")
	}

	nota := 7.89

	fmt.Printf("Sua nota será de %s porque sua nota foi %.2f \n", Conceito(nota), nota)
}
