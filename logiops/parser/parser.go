package parser

import (
	"logiops-gui/logiops"
	"strconv"
)

func Parse(s string) (logiops.LogiData, error) {
	curr := 0
	var parsedData logiops.LogiData
	parserString := lexer(s)
	err := ignoreAndDevices(&curr, &parserString, &parsedData)
	return parsedData, err
}

func ignoreAndDevices(curr *int, parserString *[]lexerTuple, parsedData *logiops.LogiData) error {
	if *curr == len(*parserString) {
		return nil
	}
	if (*parserString)[*curr].lType != keyword {
		return newParseError(*curr, "keyword", parserString)
	}

	switch (*parserString)[*curr].str {
	case "ignore":
		if parsedData.Ignore != nil {
			return newParseError(*curr, "ignore is only allowed once", parserString)
		}
		*curr++
		return ignoreParse(curr, parserString, parsedData)
	case "devices":
		if parsedData.Devices != nil {
			return newParseError(*curr, "devices is only allowed once", parserString)
		}
		*curr++
		return devicesParse(curr, parserString, parsedData)
	default:
		return newParseError(*curr, "ignore or devices", parserString)
	}
}

func ignoreParse(curr *int, parserString *[]lexerTuple, parsedData *logiops.LogiData) error {
	if (*parserString)[*curr].lType != equal {
		return newParseError(*curr, ": or =", parserString)
	}
	*curr++
	if (*parserString)[*curr].str != "[" {
		return newParseError(*curr, "[", parserString)
	}
	*curr++
	for {
		if (*parserString)[*curr].str == "]" {
			break
		} else if (*parserString)[*curr].lType == number {
			num, _ := strconv.ParseInt((*parserString)[*curr].str, 0, 64)
			parsedData.Ignore = append(parsedData.Ignore, logiops.Pid(num))
			*curr++
			if (*parserString)[*curr].lType != comma {
				break
			}
			*curr++
		} else {
			return newParseError(*curr, "PID or ]", parserString)
		}
	}
	if (*parserString)[*curr].str != "]" {
		return newParseError(*curr, "]", parserString)
	}
	*curr++
	if (*parserString)[*curr].lType == semicolon {
		*curr++
		return ignoreAndDevices(curr, parserString, parsedData)
	} else {
		return newParseError(*curr, ";", parserString)
	}

}

func devicesParse(curr *int, parserString *[]lexerTuple, parsedData *logiops.LogiData) error {
	return nil
}
