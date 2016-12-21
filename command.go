// command.go
// 2016. M.Horigome
package main

import (
	"bufio"
	"log"
	"os/exec"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// argSlicer
// "-a -b -c" string to  []string{"-a","-b","-c"} slice
func commmandOptionSlicer(line string) []string {
	s := []string{}
	item := ""
	mode := 0 // 0:Space, 1:" 2:'

	for _, c := range line {
		switch c {
		case ' ':
			if mode == 0 {
				if len(item) > 0 {
					s = append(s, item)
					item = ""
				}
			} else {
				item = item + string(c)
			}
			break

		case '"':
			if mode == 0 {
				mode = 1
			} else if mode != 1 {
				item = item + string(c)
			} else {
				s = append(s, item)
				item = ""
				mode = 0
			}
			break

		case '\'':
			if mode == 0 {
				mode = 2
			} else if mode != 2 {
				item = item + string(c)
			} else {
				s = append(s, item)
				item = ""
				mode = 0
			}
			break

		default:
			item = item + string(c)
		}
	}

	if len(item) > 0 {
		s = append(s, item)
	}

	return s
}

// shellExec
// コマンドラインを同期で実行します。実行中の出力は関数オブジェクトがCallされます。
// ※ SJISにしてます。
func commandExec(name string, option string, fn func(string)) error {

	cmd := exec.Command(name, commmandOptionSlicer(option)...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
		return err
	}
	// Stdout -> SJIS
	scan := bufio.NewScanner(transform.NewReader(stdout, japanese.ShiftJIS.NewDecoder()))
	for scan.Scan() {
		fn(scan.Text() + "\n")
	}

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
