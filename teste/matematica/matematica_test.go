package matematica

import (
	"testing"
)

//forma de fazer um teste mais unitario para cada um

// func TestMatematica(t *testing.T){
// 	valorEsperado := 10
// 	valores := Media(10, 10, 10 , 10, 10)
// 	if float64(valorEsperado) != valores{
// 		t.Errorf("A média deveria ser %.3f ao inves de %.3f", float64(valorEsperado), valores)
// 	} 
// }

func TestMatematicaDATASET(t *testing.T){
	tabela := []struct{
		valores []float64
		resultado float64
	} {
		{[]float64{10, 20 ,30}, 20},
		{[]float64{1, 2, 3}, 2},
	}


	for _, teste := range tabela{
		t.Logf("Massa: %v", teste)
		aux := Media(teste.valores...)


		if aux != teste.resultado {
			t.Errorf("Deveria ser %f mas por algo do destino, foi %f", teste.resultado, aux)
		}
	}

}