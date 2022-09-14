package parser

import (
	"strconv"
	con "logiops-gui/constants"
)



func Lexer(s string) []string {
	res := []string{}
	tmp := ""

	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case '(', ')', '{', '}', '[', ']', ';', ':', ',', '=':
			if tmp != "" {
				_, ok := con.LexerMap[tmp]
				_, err := strconv.ParseInt(tmp, 0, 64)
				if ok || err == nil {
					res = append(res, tmp)
				}
				tmp = ""
			}
			res = append(res, string(c))
		case '"':
			tmp += "\""
			i++
			for s[i] != '"' {
				tmp += string(s[i])
				i++
			}
			tmp += "\""
			res = append(res, tmp)
			tmp = ""
		case ' ', '\t', '\n':
		default:
			tmp += string(c)
		}
	}
	return res
}