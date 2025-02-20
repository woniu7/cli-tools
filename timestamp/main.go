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
			//panic(err)
			t, err := time.ParseInLocation("2006-01-02 15:04:05", os.Args[1], time.Local)
			if nil != err { panic(err) }
			fmt.Println(t.Unix())
		} else {
			fmt.Println(time.Unix(i, 0))
		}
	} else {
		fmt.Println(time.Now().Unix())
	}
}
