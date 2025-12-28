package main

import (
	"fmt"
	"strconv"
)

func Imprimir(dicionariosAprovados map[string]string) {
	fmt.Println("Lista dos alunos: ")
	dicionariosAprovados["Gustavo Senador"] = "6"
	fmt.Println(dicionariosAprovados)
}

func VerificarAprovados(dicionariosAprovados map[string]string) []string {
	listaAprovados := make([]string, 10)

	for nome, nota := range dicionariosAprovados {
		if nota, _ := strconv.Atoi(nota); nota > 6 {
			listaAprovados = append(listaAprovados, nome)
		}
	}
	return listaAprovados
}

func main(){

	dicionariosAprovados := make(map[string]string)

	dicionariosAprovados["Gustavo Senador"] = "3"
	dicionariosAprovados["Edvaldo de Faria"] = "10"
	dicionariosAprovados["Cinara Senador"] = "7"

	Imprimir(dicionariosAprovados)

	fmt.Println("Lista dos alunos aprovados: ")
	fmt.Println(VerificarAprovados(dicionariosAprovados))

	//outra forma de inicializacao
	dicionariosAprovados2 := map[string]string{
		"Gustavo Senador" : "10",
		"Edvaldo de Faria" : "10",
		"Cinara Senador" : "10", 
	}

	delete(dicionariosAprovados2,"Gustavo Senador")
	
	for nome, valor := range dicionariosAprovados2 {
		fmt.Println(nome, valor)
	}

	//map aninhado

	dicionariosAprovadosLetra := map[string]map[string]float64{
		"G" : {
			"Gustavo Senador" : 9.8,
			"Gustavo Fernandes" : 10,
		},

		"C" : {
			"Cinara Senador" : 10,
		},

		"E" : {
			"Edvaldo de Faria" : 10,
		},
	}

	dicionariosAprovadosLetra["G"]["Giuliana Granchi"] = 10.9

	for letra, valor := range dicionariosAprovadosLetra{
		fmt.Printf("\n letra %s \n", letra)
		for nome, valor := range valor {
			fmt.Println(nome, valor)
		}
	}

}
