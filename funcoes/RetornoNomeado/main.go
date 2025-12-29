package main

import (
	"fmt"
)


func dadosAlunos(alunos map[string]map[string]float32) (mediaTotal float32, alunosTotal int){
	for _, valor := range alunos{
		alunosTotal += len(valor)
		for _, nota := range valor{
			mediaTotal += nota
		}
	}
	
	mediaTotal = mediaTotal / float32(alunosTotal)

	return
}

func main(){	
	alunos := make(map[string]map[string]float32)
	alunos["C"], alunos["E"], alunos["G"] = make(map[string]float32), make(map[string]float32), make(map[string]float32) 

	alunos["G"]["Gustavo Senador"] = 9.8
	alunos["G"]["Giuliana Granchi"] = 10
	alunos["G"]["Giuliano Parreira"] = 2.3
	alunos["E"]["Edvaldo de Faria"] = 7.6
	alunos["C"]["Cinara Senador"] = 6.8

	mediaTotal, alunosTotal := dadosAlunos(alunos)

	fmt.Printf("A média total dos alunos é %.2f e o total de alunos é de %d \n", mediaTotal, alunosTotal)

}
