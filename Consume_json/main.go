package main

import (
	"fmt"
	"io/ioutil"
  	"encoding/json"
)

type Funcionarios struct {
	Fname  string `json:"fname"`
	Sname string `json:"sname"`
	Gender  string `json:"gender"`
	Age uint `json:"age"`
	Height uint    `json:"height"`
}

func main() {

  content, err := ioutil.ReadFile("dados.json")
	if err != nil {
		fmt.Println(err.Error())
	}


  var friends []Funcionarios
// json.Unmarshal(content, &friends)
  err2 := json.Unmarshal(content, &friends)
  if err2 != nil {
    fmt.Println("Error JSON Unmarshalling")
    fmt.Println(err2.Error())
  }
  for _,x := range friends{
    fmt.Printf("%s, %s, %d \n",x.Fname, x.Sname, x.Age)
  }
}