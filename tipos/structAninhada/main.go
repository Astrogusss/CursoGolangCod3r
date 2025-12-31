package main

import "fmt"

type item struct{
	produtoID string
	qnt int
	preco float64
}

type pedido struct{
	nomeID string
	itens []item
}

func (p pedido) CalcularComanda() (valorTotal float64){
	for _, item := range p.itens{
		valorTotal += item.preco * float64(item.qnt)
	}

	return
}

func (p *pedido) MudarNomeComanda(nomeAtual string){
	p.nomeID = nomeAtual
}

func main(){
	listaItens := make([]item, 0)
	pedido1 := pedido{"Gustavo Senador", listaItens}
	
	item1 := item{"Carrinho", 5, 10.5}
	item2 := item{"Controle Remoto", 1, 30}
	item3 := item{"Computador", 1, 5000}
	
	pedido1.itens = append(pedido1.itens, item1, item2, item3)

	// //calcular o preco total
	// var valorTotal float64
	// for _, pedido := range pedido1.itens{
	// 	valorTotal += pedido.preco * float64(pedido.qnt)
	// }

	fmt.Println(pedido1.CalcularComanda(), pedido1.nomeID)
	pedido1.MudarNomeComanda("Cinara Senador")
	fmt.Println(pedido1.CalcularComanda(), pedido1.nomeID)
}