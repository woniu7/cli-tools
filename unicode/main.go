package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if os.Args[1] == "-d" {
		s := strings.Split(strings.TrimSpace(os.Args[2]), "\\u")
		for _, v := range s[1:] {
			t, err := strconv.ParseInt("0x"+v, 0, 32)
			if nil != err { panic(err) }
			fmt.Printf("%c", t)
		}
	} else {
		for _, v := range os.Args[1] {
			fmt.Printf("\\u%04x", v)
		}
	}
	println();
}