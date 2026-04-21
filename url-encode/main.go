package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	//url.QueryEscape: for application/x-www-form-urlencoded, " " -> "+"
	//url.PathEscape: for url, " " -> "%20"
	fmt.Println(url.PathEscape(os.Args[1]))
}
