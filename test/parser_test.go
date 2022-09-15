package parser_test

import (
	"github.com/google/go-cmp/cmp"
	"logiops-gui/logiops"
	"logiops-gui/logiops/parser"
	"testing"
)

func TestParser(t *testing.T) {
	a := logiops.LogiData{}
	a.Ignore = []logiops.Pid{
		0xa1, 12, 0x999,
	}
	a.Devices = []logiops.LogiDevice{
		{
			Name: "MX Master",
			Dpi: 1500,
		},
		{
			Name: "MX Master 2",
		},
	}

	str := "ignore:\n[0xa1, 12, 0x999];\ndevices: ({name:\"MX Master\";dpi= 1500;}, {name= \"MX Master 2\";});"

	b, err := parser.Parse(str)

	if err != nil {
		t.Fatalf(err.Error())
	}
	if !cmp.Equal(a, b) {
		t.Fatalf("a and b are unequal.\na:%v\nb:%v", a, b)
	}
}
