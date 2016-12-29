// handler_test.go
// 2016. M.Horigome
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// /version  TEST
func Test_handlerVersion(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handlerVersion))
	defer ts.Close()

	_, err := http.Get(ts.URL + "/version")
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
}

// /command TEST
func Test_handlerCommand(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handlerCommand))
	defer ts.Close()

	r, err := http.Get(ts.URL + "/command")
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	if r.StatusCode == http.StatusOK {
		t.Fatalf("/command is POST Method only. %v", err)
	}
}
