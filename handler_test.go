// handler_test.go
// 2016. M.Horigome
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

	if r, err := http.Get(ts.URL + "/command"); err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	} else {
		if r.StatusCode == http.StatusOK {
			t.Fatalf("/command is POST Method only. %v", err)
		}
	}

	// POST test
	s1 := commands{Commands: []command{
		{Name: "echo", Option: "test"},
	}}
	b, _ := json.Marshal(s1)

	if r, err := http.NewRequest("POST", ts.URL+"/command", bytes.NewBuffer(b)); err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	} else {
		client := &http.Client{}
		rsp, _ := client.Do(r)

		if rsp.StatusCode != http.StatusOK {
			t.Fatalf("/command POST failed. %v", err)
		}

		recvBody, e := ioutil.ReadAll(rsp.Body)
		if e != nil {
			t.Fatalf("/comand response body read failed. %v", e)
		}
		if string(recvBody) != "test\n" {
			t.Fatalf("/command response result unmatch %s != %s ", string(recvBody), "test\n")
		}
	}
}
