package constants

// A map which safes the allowed keywords
var LexerMap = map[string]struct{}{
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

type Keywords int

const (
	Devices Keywords = iota
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
