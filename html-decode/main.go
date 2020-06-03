package main

import (
	"fmt"
	"html"
	"os"
)

func main() {
	fmt.Println(html.UnescapeString(os.Args[1]))
}
