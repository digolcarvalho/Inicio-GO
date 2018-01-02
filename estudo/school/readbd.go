package main

import (
	"fmt"
  _ "github.com/go-sql-driver/mysql"
	"database/sql"
	"net/http"
)

var db, err = sql.Open("mysql", "root:abcd1234@/dbgolang?charset=utf8")
// estou criando 2 variaveis para usar com o meu banco e passando os dados de acesso.

type Qry struct{// é necessário a primeira maiuscula, senão eu não consigo acessar a informação
	id int
	fname string
	sname string
	email string
}

func main(){
	
	rows, err := db.Query("Select * from user")
	checkError(err)

	itens := []Qry{}
	
	for rows.Next(){
		qry := Qry{}
		rows.Scan(&qry.id, &qry.sname, &qry.fname, &qry.email)
		itens = append(itens, qry)
		}


fmt.Println(http.ListenAndServe(":8080", nil))

	
}

func checkError(err error){ // função para checar os erros nas funções
	if err != nil{
		panic(err)
	}
}