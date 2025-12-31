package main

import (
	"encoding/json"
	"fmt"
)

type Estudante struct{
	Nome string `json:"id"`
	Sobrenome string `json:"sobrenome"`
	Nota float32 `json:"nota`
}

func main(){	
	//struct para um json
	aluno1 := Estudante{"Gustavo", "Senador", 9.8}
	aluno1JSON, _ := json.Marshal(aluno1)

	fmt.Println(string(aluno1JSON))

	//json para uma struct

	var aluno2 Estudante
	aluno2JSON := `{"id":"Cinara","sobrenome":"Senador","nota":10}`
	json.Unmarshal([]byte(aluno2JSON), &aluno2)

	fmt.Println(aluno2)

}