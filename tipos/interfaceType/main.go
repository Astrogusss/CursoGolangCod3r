package main

import (
	"fmt"
	"reflect"
)

func Imprimir(coisa1 interface{}){
	fmt.Println(reflect.TypeOf(coisa1))
}


func main(){
	//meio que fodase
	var coisa interface{} = 20
	fmt.Println(coisa)

	coisa = "Vai tomar no cu"
	fmt.Println(coisa)

	Imprimir("ifhsdh")

}