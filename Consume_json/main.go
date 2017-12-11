package main

import (
	"fmt"
	"io/ioutil"
  	"encoding/json"
)

type Friend struct {
	Fname  string `json:"fname"`
	Sname string `json:"sname"`
	Gender  string `json:"gender"`
	Height int    `json:"height"`
}

func main() {

  content, err := ioutil.ReadFile("friends.json")
	if err != nil {
		fmt.Println(err.Error())
	}


  var friends []Friend
// json.Unmarshal(content, &friends)
  err2 := json.Unmarshal(content, &friends)
  if err2 != nil {
    fmt.Println("Error JSON Unmarshalling")
    fmt.Println(err2.Error())
  }
  for _,x := range friends{
    fmt.Printf("%s \n",x.Fname)
  }



}
}