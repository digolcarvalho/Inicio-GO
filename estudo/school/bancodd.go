package main

import (
	"fmt"
  _ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var db, err = sql.Open("mysql", "root:abcd1234@/dbgolang")
// estou criando 2 variaveis para usar com o meu banco e passando os dados de acesso.

func main(){
	
	// db.Ping()
	
	// stmt, err := db.Prepare("Insert into user (fname, sname, email) values (?, ?, ?)")
	// _ , err = stmt.Exec("Isabella", "Carvalho", "email teste 3")
	// checkError(err)

	//db.Close()


}

func checkError(err error){ // função para checar os erros nas funções
	if err != nil{
		panic(err)
	}else{
		fmt.Println("Dados inseridos com sucesso.")
	}
}