package main

import (
	"fmt"
    "github.com/gorilla/mux"
	"database/sql"
	"net/http"
	"html/template"
	_"github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "root:abcd1234@/dbgolang")
// estou criando 2 variaveis para usar com o meu banco e passando os dados de acesso.

type Qry struct{
	Id int
	Fname string
	Sname string
	Email string
}

func main(){

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	fmt.Println(http.ListenAndServe(":8080", r))
}
	

func HomeHandler(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("index.html"))
	 	 if err := t.ExecuteTemplate(w, "index.html", SelectItens()); err!=nil { // Não podemos esquecer de colocar a variável com os valores 
	 		http.Error(w, err.Error(), http.StatusInternalServerError) //nesse caso é o slice itens, com todas as informações inserridas.
	 	 }
	  }
	

func SelectItens() []Qry {

	rows, err := db.Query("Select * from user")
	checkError(err)

	itens := []Qry{}
	
	for rows.Next(){
		qry := Qry{ }
		rows.Scan(&qry.Id,&qry.Fname, &qry.Sname, &qry.Email)
		itens = append(itens, qry)
		//fmt.Println(itens[len(itens)-1])
		}
		return itens
}

func checkError(err error){ // função para checar os erros nas funções
	if err != nil{
		panic(err)
	}
}