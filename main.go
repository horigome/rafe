// main.go
// 2016. M.Horigome
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"runtime"

	"net/http"
)

// command usage
var commandUsage = `
rafe api service

usage:
    rafe <option>
option:
`

// options
type options struct {
	portNo int
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
func handlerVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("! method error.")
		return
	}
	v := "{ version: \"1.0.0.0\" }"
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, v)
}

// handlerCommand
func handlerCommand(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Println("! method error.")
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("! Read body err.", err)
		return
	}
	fmt.Println(string(content))
	var cmd commands
	err = json.Unmarshal(content, &cmd)
	if err != nil {
		fmt.Println("! Json Unmarshal err. ", err)
		return
	}
	fmt.Println("Recv OK")

	// debug --
	cmd.Print()
	// --

	ret := ""
	for _, exec := range cmd.Commands {
		err := commandExec(exec.Name, exec.Option, func(out string) {
			ret += out
		})
		if err != nil {
			fmt.Println("! exec err. ", err)
		}
	}
	w.Header().Set("Content-Type", "text")
	fmt.Fprint(w, ret)
}

// makeOptions
func makeOptions() options {

	var o options
	flag.IntVar(&o.portNo, "port", 8080, " listen port no ")

	flag.Usage = func() {
		fmt.Println(commandUsage)
		flag.PrintDefaults()
	}

	flag.Parse()
	return o
}

// main
func main() {

	opt := makeOptions()
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	http.HandleFunc("/command", handlerCommand)
	http.HandleFunc("/version", handlerVersion)

	fmt.Println("==> start server.")
	http.ListenAndServe(fmt.Sprintf(":%d", opt.portNo), nil)
}
