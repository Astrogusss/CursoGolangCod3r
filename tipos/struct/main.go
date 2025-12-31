package main

import "fmt"

type pessoa struct{
	nome string
	idade int8
}

type produto struct{
	nome string
	preco float64
	desconto float64
}

//funcao com recptor (receiver) 
//funciona mais como um metodo em OO
func (p produto) precoComDesconto() float64{
	return p.preco * (1 - p.desconto)
}

func imprimirLista(lista []*pessoa){
	for _, pessoa := range lista{
		fmt.Printf("O nome do indivíduo é %s e sua idade é %d \n", pessoa.nome, pessoa.idade)
	}
}

func main(){
	vetorPessoas := make([]*pessoa, 0)
	vetorPessoas = append(vetorPessoas, &pessoa{"Gustavo Senador", 20})
	vetorPessoas = append(vetorPessoas, &pessoa{"Cinara Senador", 48})
	vetorPessoas = append(vetorPessoas, &pessoa{"Edvaldo de Faria", 51})
	imprimirLista(vetorPessoas)
	
	
	//metodos de acessar propriedades da propria struct
	produto1 := produto{"Carrinho", 100, 0.5}
	fmt.Println(produto1, produto1.precoComDesconto())
}