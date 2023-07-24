package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func encodeHandler(w http.ResponseWriter, r *http.Request){
	peter := User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:	   25,
	}

	json.NewEncoder(w).Encode(peter)  // peterの情報をJSONにencode
}

func decodeHandler(w http.ResponseWriter, r *http.Request){
	var 
}


func main() {
	http.HandleFunc("/encode", encodeHandler)
	http.HandleFunc("/decode", decodeHandler)

	http.ListenAndServe(":8080", nil)
}