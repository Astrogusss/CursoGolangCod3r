package main

import (
	"fmt"
	"math"
)

type area interface{
	calculo() float64
	AlterarRaio()
}

type circulo struct{
	raio float64
}
type quadrado struct{
	lado float64
}
type retangulo struct{
	base float64
	largura float64
}

func (c circulo) calculo() float64{
	return math.Pow(c.raio, 2) * math.Pi
}
func (q quadrado) calculo() float64{
	return math.Pow(q.lado, 2)
}
func (r retangulo) calculo() float64{
	return r.base * r.largura
}

//a interface é implementada de forma implicita
func ImprimirArea (a area) {
	fmt.Printf("%.2f \n",a.calculo())
}

//funcao para alterar o raio do circulo
func (c *circulo) AlterarRaio(){
	c.raio = 10
}

func main(){
	bolinha := circulo{3}
	fmt.Printf("Imprimir a área da bolinha (circulo) antes: ")
	ImprimirArea(&bolinha)

	bolinha.AlterarRaio()
	fmt.Printf("Imprimir a área da bolinha (circulo) depois: ")
	ImprimirArea(&bolinha)

	// quadrado := quadrado{3}
	// fmt.Printf("Imprimir a área do quadrado ")
	// ImprimirArea(quadrado)

	// retangulo := retangulo{3, 5}
	// fmt.Printf("Imprimir a área do retangulo ")
	// ImprimirArea(retangulo)


	var circuloBonitinho area = &circulo{10}
	circuloBonitinho.AlterarRaio()
	ImprimirArea(circuloBonitinho)
}