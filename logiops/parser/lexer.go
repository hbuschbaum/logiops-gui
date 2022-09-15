package parser

import (
	con "logiops-gui/logiops/constants"
	"strconv"
)

func Lexer(s string) []string {
	res := []string{}
	tmp := ""

	for i := 0; i < len(s); i++ { //Iterate over the string
		switch c := s[i]; c {
		case '(', ')', '{', '}', '[', ']', ';', ':', ',', '=': //One of those keywords where found
			if tmp != "" { 
				_, ok := con.LexerMap[tmp] //Look if the tmp-string is a keyword
				_, err := strconv.ParseInt(tmp, 0, 64) //Look if the tmp-string is a number
				if ok || err == nil {
					res = append(res, tmp) //append the tmp-string
				}
				tmp = ""
			}
			res = append(res, string(c)) //append the current rune
		case '"': //If there is a " look for the closing " and put the whole string (incl. "") in res
			tmp += "\""
			i++
			for s[i] != '"' {
				tmp += string(s[i])
				i++
			}
			tmp += "\""
			res = append(res, tmp)
			tmp = ""
		case ' ', '\t', '\n': //If whitespace don't do anything
		default:
			tmp += string(c) 
		}
	}
	return res
}
