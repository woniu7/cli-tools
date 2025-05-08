package main

import (
	"fmt"
	"os"
	"strings"
)

// stringToHex converts a string to \x format (e.g., "Hello" -> "\x48\x65\x6c\x6c\x6f")
func stringToHex(input string) string {
	var builder strings.Builder
	for _, char := range []byte(input) {
		builder.WriteString(fmt.Sprintf("\\x%02x", char))
	}
	return builder.String()
}

// just echo -e "\x48\x65\x6c\x6c\x6f"
// hexToString converts a \x formatted hex string (e.g., "\x48\x65\x6c\x6c\x6f") to a string (e.g., "Hello")
//func hexToString(input string) (string, error) {
//}

func main() {
	input := os.Args[1]
	// 转换为 \x 格式
	result := stringToHex(input)
	fmt.Printf("Input: %s\nOutput: %s\n", input, result)
}
