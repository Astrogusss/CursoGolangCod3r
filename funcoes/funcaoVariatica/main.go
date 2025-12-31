package main

import "fmt"

//ele vem como vetor
func media(numeros ...float64) (resultado float64){
	for _, valor := range numeros{
		resultado += valor
	}

	resultado /= float64(len(numeros))
	return
}


func imprimirCandidatos(aprovados ...string){
	for _, nome := range aprovados{
		fmt.Println(nome)
	}
}

func main(){
	fmt.Println(media(1.1, 32.1, 10, 123, 13))

	//assim que ce passa um slice por referencia
	candidatos := []string{"Gustavo Senador", "Edvaldo de Faria", "Cinara Senador", "Giuliana Granchi"}
	imprimirCandidatos(candidatos...)
}