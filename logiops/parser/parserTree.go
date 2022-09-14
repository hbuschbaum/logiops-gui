package parser

import (
	"logiops-gui/helper"
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

func MakeParseTree(s []string) (Tree, error) {
	children, err := ignoreAndDevices(s)
	if err != nil {
		return Tree{}, err
	}
	return Tree{children}, nil
}

var curr = 0
var parserString string

//TODO make a real parser... i.e. go from position 0 to end and not from keyword to keyword.

func ignoreAndDevices(s []string) ([]Node, error) {
	res := []Node{}
	if ignoreIndex, ignoreFound := helper.Index(s, "ignore"); ignoreFound {
		if (s[ignoreIndex+1] == ":" || s[ignoreIndex+1] == "=") && s[ignoreIndex+2] == "[" {
			ignoreClose, closeFound := helper.Index(s[ignoreIndex+2:], "]")
			if !closeFound || s[ignoreClose+1] != ";" {
				return nil, helper.ParseError{Err: "Expected '];'"}
			}
			n := NewNode("ignore")
			children, err := ignoreParse(s[ignoreIndex+2 : ignoreClose+1])
			if err != nil {
				return nil, err
			}
			n.AppendAll(children)
			res = append(res, n)
		} else {
			return nil, helper.ParseError{Err: "Expected ':' or '=' after ignore followed by '["}
		}
	}

	if devicesIndex, devicesFound := helper.Index(s, "devices"); devicesFound {
		if (s[devicesIndex+1] == ":" || s[devicesIndex+1] == "=") && s[devicesIndex+2] == "(" {
			devicesClose, closeFound := helper.LastIndex(s[devicesIndex+2:], ")")
			if !closeFound || s[devicesClose+1] != ";" {
				return nil, helper.ParseError{Err: "Expected ');'"}
			}

			n := NewNode("devices")
			children, err := devicesParse(s[devicesIndex+2 : devicesClose+1])
			if err != nil {
				return nil, err
			}
			n.AppendAll(children)
			res = append(res, n)
		}
	}

	return res, nil
}

func ignoreParse(s []string) ([]Node, error) {
	return nil, nil
}

func devicesParse(s []string) ([]Node, error)
