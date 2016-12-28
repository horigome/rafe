// main.go
// 2016. M.Horigome
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"net/http"
)

var version = "0.0.0.0"

// command usage
var commandUsage = `
rafe api service

usage:
    rafe <option>
option:
`

// options
type options struct {
	portNo int    // Service Port (default 8080)
	locale string // "UTF8(Default)", or "SJIS"
}

// optionsGlobal Command options (global)
var optionsGlobal = options{}

// initOptions
func initOptions() {

	var v bool

	flag.IntVar(&optionsGlobal.portNo, "port", 8080, "Listen port")
	flag.StringVar(&optionsGlobal.locale, "locale", "UTF8", "Locale")
	flag.BoolVar(&v, "version", false, "Version.")

	flag.Usage = func() {
		fmt.Println(commandUsage)
		flag.PrintDefaults()
	}

	flag.Parse()

	if v == true {
		fmt.Println("rafe. version ", version)
		os.Exit(0)
	}
}

// main
func main() {

	initOptions()
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	// http handler
	http.HandleFunc("/command", handlerCommand)
	http.HandleFunc("/version", handlerVersion)

	// start
	fmt.Println("==> start server. http://localhost:", optionsGlobal.portNo)
	http.ListenAndServe(fmt.Sprintf(":%d", optionsGlobal.portNo), nil)
}
