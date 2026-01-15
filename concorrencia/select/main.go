package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

//select funciona mais como um if para canais
func OmaisRapido(url1, url2, url3 string) string{
	canal1 := Titulos(url1)
	canal2 := Titulos(url2)
	canal3 := Titulos(url3)

	select{
	case t1 := <-canal1:
		return t1
	case t2 := <-canal2:
		return t2
	case t3 := <-canal3:
		return t3
	}
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

	fmt.Println(OmaisRapido("https://unifei.edu.br", "https://www.udemy.com", "https://www.linkedin.com"))
}