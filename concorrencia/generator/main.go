package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

//aqui estamos dizendo que irá retornar somente um canal de leitura, não poderá ser colocado nada no canal
func Titulos(sites ...string) <-chan string {
	c := make(chan string)

	for _, url := range sites {
		go func(url string){
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			if peidao := r.FindStringSubmatch(string(html)); len(peidao) > 1 { 
				c <- peidao[1]
			} else {
				c <- "Error"
			}
		}(url)
	}
	return c
}

func main(){
	titulos1 := Titulos("https://unifei.edu.br", "https://www.facebook.com")
	titulo2 := Titulos("https://www.linkedin.com", "https://github.com")

	fmt.Println("primeiro 1: ", <-titulos1, "primeiro 2: ", <-titulo2)
	fmt.Println("Segundo 1: ", <-titulos1, "Segundo 2: ", <-titulo2)
}