package main

import "fmt"


func ParOuImpar(num int8) string{
	//ele atrasa a sua execucao até a ultima linha, só será executar no final de toda execucção da função
	//muito utilizado para fechar recursos como banco de dados, fechar arquivos, etc -> ao inves de ficar fazendo o mesmo script para cada return, pode so usar o defer que será bem melhor
	defer fmt.Println("Final da função")
	fmt.Println("Teste")

	if num % 2 == 0 {
		return "par"
	} else {
		return "impar"
	}
}

func main(){
	fmt.Printf("o seu número é %s \n", ParOuImpar(27))
}