package main

import (
	"fmt"
	"html"
	"os"
)

func main() {
	fmt.Println(html.EscapeString(os.Args[1]))
}
