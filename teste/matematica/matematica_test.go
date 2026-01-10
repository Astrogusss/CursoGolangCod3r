package matematica

import (
	"testing"
)

func TestMatematica(t *testing.T){
	valorEsperado := 9
	valores := Media(10, 10, 10 , 10, 10)

	if float64(valorEsperado) != valores{
		t.Errorf("A média deveria ser %.3f ao inves de %.3f", float64(valorEsperado), valores)
	} 
}