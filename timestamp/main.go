package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		i, err := strconv.ParseInt(os.Args[1], 10, 64)
		if nil != err {
			panic(err)
		}
		fmt.Println(time.Unix(i, 0))
	} else {
		fmt.Println(time.Now().Unix())
	}
}
