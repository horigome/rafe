// handler.go
// 2016. M.Horigome
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// version type
type versionStruct struct {
	Version     string `json:"version"`
	Description string `json:"description"`
}

// command type
type command struct {
	Name   string `json:"name"`
	Option string `json:"option"`
}

// commands type
type commands struct {
	Commands []command `json:"commands"`
}

// Print method
func (p *commands) Print() {
	// debug
	for i, v := range p.Commands {
		fmt.Printf("Index[ %d ]\n", i)
		fmt.Println(" name   :", v.Name)
		fmt.Println(" option :", v.Option)
		fmt.Println("----------------------")
	}
}

// handlerVersion
// [GET] http://host:port/version
func handlerVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("! method error.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	v := versionStruct{Version: "1.0.0.0", Description: "rafe command service"}
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println("! json marshal error. > ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// handlerCommand
// [POST] http://host:port/command
func handlerCommand(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Println("! method error.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("! Read body err.", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(string(content))
	var cmd commands
	err = json.Unmarshal(content, &cmd)
	if err != nil {
		fmt.Println("! Json Unmarshal err. ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Recv OK")

	// debug --
	cmd.Print()
	// --
	w.Header().Set("Content-Type", "text")

	ret := ""
	for _, exec := range cmd.Commands {
		err := commandExec(exec.Name, exec.Option, func(out string) {
			ret += out
		})
		if err != nil {
			fmt.Println("! exec err. ", err)
		}
		fmt.Fprint(w, ret)
	}

}
