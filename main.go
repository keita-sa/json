package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	peter := User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}

	json.NewEncoder(w).Encode(peter) // peterの情報をJSONにencode
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	var user User                         // structで定義されたUser型の変数user
	json.NewDecoder(r.Body).Decode(&user) // JSONデータを読み込み、decodeして構造体Userのuserのアドレスに格納

	fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age) // %sは文字列、%dは整数
}

func main() {
	http.HandleFunc("/encode", encodeHandler)
	http.HandleFunc("/decode", decodeHandler)

	http.ListenAndServe(":8080", nil)
}
