// handler_test.go
// 2016. M.Horigome
package main

import (
	"testing"
)

func Test_commmandOptionSlicer(t *testing.T) {

	// case.1
	s1Ans := []string{"a", "b", "cc", "d"}

	for i, v := range commmandOptionSlicer("a b cc d") {
		if v != s1Ans[i] {
			t.Fatalf("anmatch slice. Idx:%d , %s(in) != %s(ref)", i, v, s1Ans[i])
		}
	}

	// case.2
	s2Ans := []string{"hello", "wor ld", "!"}

	for i, v := range commmandOptionSlicer("hello 'wor ld' !") {
		if v != s2Ans[i] {
			t.Fatalf("anmatch slice. Idx:%d , %s(in) != %s(ref)", i, v, s2Ans[i])
		}
	}

	// case.3
	s3Ans := []string{"he llo", "wor ld", "!"}

	for i, v := range commmandOptionSlicer("\"he llo\" 'wor ld'  !") {
		if v != s3Ans[i] {
			t.Fatalf("anmatch slice. Idx:%d , %s(in) != %s(ref)", i, v, s3Ans[i])
		}
	}
}

func Test_commmandExec(t *testing.T) {

	var out string
	e := commandExec("echo", "test", func(s string) {
		out += s
	})
	if e != nil {
		t.Fatalf("commandExec failed. %s\n", e)
	}
	if out != "test\n" {
		t.Fatalf("commandExec return failed. %s (in) != %s (result)", "test", out)
	}
}
