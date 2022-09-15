package parser
import (
	"fmt"
)

/// Lexer Struct and Constants -------------------------------------------------------
type lexerType int

const (
	equal lexerType = iota
	parentheses
	semicolon
	comma
	keyword
	number
	strings
)

type lexerTuple struct {
	lType lexerType
	str   string
}

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

type keywords int

const (
	Devices keywords = iota
	Name
	Ignore
	True
	False
	Cid
	Action
	Type
	Gestures
	Direction
	Mode
	Keys
	Buttons
	Smartshift
	On
	Threshold
	Hiresscroll
	Hires
	Invert
	Target
	Dpi
)


/// Parse Error struct --------------------------------------------------------------
type ParseError struct {
	Pos      int
	Expected string
	Wrong    string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("Pos %v: %v. Expected: %v", e.Pos, e.Wrong, e.Expected)
}

func newParseError(pos int, expected string, wrong *[]lexerTuple) ParseError {
	return ParseError{Pos: pos, Expected: expected, Wrong: (*wrong)[pos].str}
}