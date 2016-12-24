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

// version type
type version struct {
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
	v := version{Version: "1.0.0.0", Description: "rafe command service"}
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

	// http handler
	http.HandleFunc("/command", handlerCommand)
	http.HandleFunc("/version", handlerVersion)

	// start
	fmt.Println("==> start server. http://localhost:", opt.portNo)
	http.ListenAndServe(fmt.Sprintf(":%d", opt.portNo), nil)
}
