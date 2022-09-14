package parser

import (
	"strconv"
)

// A map which safes the allowed keywords
var lexerMap = map[string]struct{}{
	"devices":     {},
	"name":        {},
	"ignore":      {},
	"true":        {},
	"false":       {},
	"cid":         {},
	"action":      {},
	"type":        {},
	"gestures":    {},
	"direction":   {},
	"mode":        {},
	"keys":        {},
	"buttons":     {},
	"smartshift":  {},
	"on":          {},
	"threshold":   {},
	"hiresscroll": {},
	"hires":       {},
	"invert":      {},
	"target":      {},
	"dpi":         {},
}

func Lexer(s string) []string {
	res := []string{}
	tmp := ""

	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case '(', ')', '{', '}', '[', ']', ';', ':', ',':
			if tmp != "" {
				_, ok := lexerMap[tmp]
				_, err := strconv.Atoi(tmp)
				if ok || err != nil {
					res = append(res, tmp)
					tmp = ""
				}
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