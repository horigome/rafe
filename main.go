// main.go
// 2016. M.Horigome
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type command struct {
	Name   string `json:"name"`
	Option string `json:"option"`
}

type commands struct {
	Commands []command `json:"commands"`
}

func handlerCommand(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err1.", err)
		return
	}
	fmt.Println(string(content))
	var cmd commands
	err = json.Unmarshal(content, &cmd)
	if err != nil {
		fmt.Println("err2.", err)
		return
	}
	fmt.Println("OK")
	// debug
	for _, v := range cmd.Commands {
		fmt.Println("name   :", v.Name)
		fmt.Println("option :", v.Option)
	}
}

func handlerVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	v := "{ version: \"1.0.0.0\" }"
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, v)
}

func main() {
	http.HandleFunc("/command", handlerCommand)
	http.HandleFunc("/version", handlerVersion)
	http.ListenAndServe(":8080", nil)
}
