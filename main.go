package main

import (
	"fmt"
	"logiops-gui/src/parser"
)

func main() {
	test := parser.Lexer("devices: ({name: \"Hello World\";},\n{name: \"Testi test test\";}); ignore: [100]")
	for _, t := range test {
		fmt.Println(t)
	}
}
