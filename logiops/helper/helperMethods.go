package helper

import "fmt"

func Index(s []string, str string) (int, bool) {
	for i := range s {
		if s[i] == str {
			return i, true
		}
	}

	return -1, false
}

func LastIndex(s []string, str string) (int, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == str {
			return i, true
		}
	}

	return -1, false
}

type ParseError struct {
	Err string
}

func (e ParseError) Error() string {
	return fmt.Sprint(e.Err)
}
