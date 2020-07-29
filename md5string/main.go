package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	h := md5.Sum([]byte(os.Args[1]))
	fmt.Printf("%x\n", h)
	// fmt.Println(hex.EncodeToString(h[:]))
}
