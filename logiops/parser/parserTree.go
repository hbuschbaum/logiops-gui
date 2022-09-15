package parser

import (
	"logiops-gui/logiops/helper"
	"logiops-gui/logiops"
	"strconv"
)

type Node struct {
	Children []Node
	Value    string
}

type Tree struct {
	Children []Node
}

func (n *Node) Append(newNode Node) {
	n.Children = append(n.Children, newNode)
}

func (n *Node) AppendAll(newNodes []Node) {
	n.Children = append(n.Children, newNodes...)
}

func NewNode(val string) Node {
	return Node{Value: val}
}

// func MakeParseTree(s []string) (Tree, error) {
// 	children, err := ignoreAndDevices(s)
// 	if err != nil {
// 		return Tree{}, err
// 	}
// 	return Tree{children}, nil
// }

var curr = 0
var parserString []string
var parsedData logiops.LogiData 


func Parse(s []string) (logiops.LogiData, error) {
	parserString = s
	err := ignoreAndDevices()
	return parsedData, err
}

func ignoreAndDevices() error {
	if curr == len(parserString) {
		return nil
	}

	switch parserString[curr] {
	case "ignore":
		curr++
		return ignoreParse()
	case "devices":
		curr++
		return devicesParse()
	default:
		return helper.ParseError{Pos: curr, Expected: "ignore or devices" ,Wrong: parserString[curr]}
	}
}

func ignoreParse() error {
	if !(parserString[curr] == ":" || parserString[curr] == "=") {
		return helper.ParseError{Pos: curr, Expected: ": or =", Wrong: parserString[curr]}
	}
	curr++
	if parserString[curr] != "[" {
		return helper.ParseError{Pos: curr, Expected: "[", Wrong: parserString[curr]}
	}
	curr++
	for {
		if parserString[curr] == "]" {
			break
		} else if num, err := strconv.ParseInt(parserString[curr], 0, 64); err == nil {
			parsedData.Ignore = append(parsedData.Ignore, logiops.Pid(num))
			curr++
			if parserString[curr] != "," {
				break
			}
			curr++
		} else {
			return helper.ParseError{Pos: curr, Expected: "Pid or ]", Wrong: parserString[curr]}
		}
	}
	if parserString[curr] != "]" {
		return helper.ParseError{Pos: curr, Expected: "]", Wrong: parserString[curr]}
	}
	curr++
	if parserString[curr] == ";" {
		curr++
		return ignoreAndDevices()
	} else {
		return helper.ParseError{Pos: curr, Expected: "}", Wrong: parserString[curr]}
	}
	
}

func devicesParse() error {
	return nil
}

//TODO make a real parser... i.e. go from position 0 to end and not from keyword to keyword.

// func ignoreAndDevices(s []string) ([]Node, error) {
// 	res := []Node{}
// 	if ignoreIndex, ignoreFound := helper.Index(s, "ignore"); ignoreFound {
// 		if (s[ignoreIndex+1] == ":" || s[ignoreIndex+1] == "=") && s[ignoreIndex+2] == "[" {
// 			ignoreClose, closeFound := helper.Index(s[ignoreIndex+2:], "]")
// 			if !closeFound || s[ignoreClose+1] != ";" {
// 				return nil, helper.ParseError{Err: "Expected '];'"}
// 			}
// 			n := NewNode("ignore")
// 			children, err := ignoreParse(s[ignoreIndex+2 : ignoreClose+1])
// 			if err != nil {
// 				return nil, err
// 			}
// 			n.AppendAll(children)
// 			res = append(res, n)
// 		} else {
// 			return nil, helper.ParseError{Err: "Expected ':' or '=' after ignore followed by '["}
// 		}
// 	}

// 	if devicesIndex, devicesFound := helper.Index(s, "devices"); devicesFound {
// 		if (s[devicesIndex+1] == ":" || s[devicesIndex+1] == "=") && s[devicesIndex+2] == "(" {
// 			devicesClose, closeFound := helper.LastIndex(s[devicesIndex+2:], ")")
// 			if !closeFound || s[devicesClose+1] != ";" {
// 				return nil, helper.ParseError{Err: "Expected ');'"}
// 			}

// 			n := NewNode("devices")
// 			children, err := devicesParse(s[devicesIndex+2 : devicesClose+1])
// 			if err != nil {
// 				return nil, err
// 			}
// 			n.AppendAll(children)
// 			res = append(res, n)
// 		}
// 	}

// 	return res, nil
// }

// func ignoreParse(s []string) ([]Node, error) {
// 	return nil, nil
// }

// func devicesParse(s []string) ([]Node, error)
