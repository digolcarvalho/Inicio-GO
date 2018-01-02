package main

import (
	"fmt"
	"net/http"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Post struct {
	Id int
	Title string
	Body string
}

func main(){

	//Criando uma resposta simples e direta do servidor
		//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		//fmt.Fprintf(w, "Olá Mundo")
	//})

	/*Criando uma func para o carregamento de arquivos html
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			 
			post := Post {Id: 1, Title: "Sucesso", Body: "teste"}			

			if uptitle := r.FormValue("uptitle"); uptitle != "" {
				post.Title = uptitle
			}	
			
			if upbody := r.FormValue("upbody"); upbody != "" {
				post.Body = upbody
			}	
			
			t := template.Must(template.ParseFiles("index.html"))
	 		if err := t.ExecuteTemplate(w, "index.html", post); err!=nil { // Não podemos esquecer de colocar a variável com os valores (post)
				http.Error(w, err.Error(), http.StatusInternalServerError)
		 	}
	 	})
	
	fmt.Println(http.ListenAndServe(":8080", nil))
	*/
}

