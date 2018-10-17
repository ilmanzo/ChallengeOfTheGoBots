// program to be profiled

package main

import (
	"fmt"
	"strings"
)

//slow implementation, just for demo
func leftpad1(s string, length int, char rune) string {
	for len(s) < length {
		s = string(char) + s
	}
	return s
}

//a better one
func leftpad2(s string, length int, char rune) string {
	if len(s) < length {
		return strings.Repeat(string(char), length-len(s)) + s
	}
	return s
}

func main() {

	fmt.Println(leftpad1("ciao!", 10, '.'))

}
