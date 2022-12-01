package main

import (
	"fmt"
	"os"
	"strings"
)

func (d Dir) GoFileCheck() {
	fmt.Println("Go File Check is Starting...")
	files, err := os.ReadDir(string(d))
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Go File Found."
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("File: ", gofile, " ", text)
	}
}

func (sd StrDir) GoFileCheck() {
	fmt.Println("Go File Check is Starting...")
	files, err := os.ReadDir(sd.strDirect)
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Go File Found."
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("File: ", gofile, " ", text)
	}
}
