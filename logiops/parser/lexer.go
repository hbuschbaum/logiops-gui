package parser

import (
	"strconv"
)

func lexer(s string) []lexerTuple {
	res := []lexerTuple{}
	tmp := ""

	for i := 0; i < len(s); i++ { //Iterate over the string
		switch c := s[i]; c {
		case '(', ')', '{', '}', '[', ']', ';', ':', ',', '=': //One of those keywords where found
			if tmp != "" {
				if _, ok := lexerMap[tmp]; ok { //Look if the tmp-string is a keyword
					res = append(res, lexerTuple{keyword, tmp}) //append the tmp-string
				} else if _, err := strconv.ParseInt(tmp, 0, 64); err == nil { //Look if the tmp-string is a number
					res = append(res, lexerTuple{number, tmp})
				}
				tmp = ""
			}
			switch c {
			case '(', ')', '}', '{', '[', ']':
				res = append(res, lexerTuple{parentheses, string(c)}) //append the current rune
			case ';':
				res = append(res, lexerTuple{semicolon, string(c)}) //append the current rune
			case ':', '=':
				res = append(res, lexerTuple{equal, string(c)}) //append the current rune
			default:
				res = append(res, lexerTuple{comma, string(c)}) //append the current rune
			}

		case '"': //If there is a " look for the closing " and put the whole string (incl. "") in res
			tmp += "\""
			i++
			for s[i] != '"' {
				tmp += string(s[i])
				i++
			}
			tmp += "\""
			res = append(res, lexerTuple{strings, tmp})
			tmp = ""
		case ' ', '\t', '\n': //If whitespace don't do anything
		default:
			tmp += string(c)
		}
	}
	return res
}
