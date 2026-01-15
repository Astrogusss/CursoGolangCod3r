package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

//padrao de concorrencia que juntas dois canais em somente um
func encaminhar (leitura <-chan string, destino chan string){
	for {
		destino <- <-leitura
	}
}

func juntar (entrada1, entrada2 <-chan string) <-chan string{
	resposta := make(chan string)

	go encaminhar(entrada1, resposta)
	go encaminhar(entrada2, resposta)
	
	return resposta
}

func Titulos(lista ...string) <-chan string{
	c := make(chan string)

	for _, url := range lista {
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

	canalMisto := juntar(Titulos("https://unifei.edu.br", "https://www.udemy.com"), Titulos("https://www.linkedin.com", "https://github.com"))

	for valor := range canalMisto {
		fmt.Println(valor)
	}

}