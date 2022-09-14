package parser_test

import (
	"fmt"
	con "logiops-gui/logiops/constants"
	"logiops-gui/logiops/parser"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestLexer(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	for j := 0; j < 10; j++ {
		arr := makeRandomArray()
		str := toString(arr)
		newarr := parser.Lexer(str)

		for i, v := range arr {
			if v != newarr[i] {
				t.Fatalf("Not equal on position %v: %v. It should be: %v\nWhole array: %v\n\nShould be: %v", i, newarr[i], v, newarr, arr)
			}
		}
	}
}

func toString(s []string) string {
	var sb strings.Builder
	for _, c := range s {
		sb.WriteString(c)
		for i := 0; i < rand.Int()%100; i++ {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

func makeRandomArray() []string {
	keys := make([]string, 0, len(con.LexerMap))
	for k := range con.LexerMap {
		keys = append(keys, k)
	}

	res := make([]string, 100)

	for i := 0; i < 99; i++ {
		switch rand.Int() % 3 {
		case 0:
			res[i] = keys[rand.Int()%len(keys)]
			i++
			if rand.Int()%2 == 0 {
				res[i] = ":"
			} else {
				res[i] = "="
			}
		case 1:
			res[i] = "\"Hellolololololo +da432\""
			i++
			res[i] = ";"
		case 2:
			if rand.Int()%2 == 0 {
				res[i] = fmt.Sprint(rand.Int())
			} else {
				res[i] = fmt.Sprintf("0x%x", rand.Int())
			}
			i++
			res[i] = ","

		}

	}

	return res
}
