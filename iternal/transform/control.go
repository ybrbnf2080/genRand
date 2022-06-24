package transform

import (
	"bufio"
	"os"
)

type Transform struct {
	Height       int
	Width        int
	HeightOfset  int
	WidhtOfset   int
	CompressCoof int
}

func NewTransform(compress int, height int, width int) *Transform {

	var transform = Transform{
		Height:       height,
		Width:        width,
		HeightOfset:  0,
		WidhtOfset:   0,
		CompressCoof: compress,
	}

	return &transform
}

func (t *Transform) Controler(stdin bufio.Reader) {

	b := nextByte(&stdin)
	//fmt.Println("I got the byte", b, "("+string(b)+")")
	switch string(b) {

	case "+":
		t.CompressCoof = t.CompressCoof - 1
	case "-":
		t.CompressCoof = t.CompressCoof + 1
	case "q":
		os.Exit(1)
	case "l":
		t.HeightOfset = t.HeightOfset + 1
	case "h":
		t.HeightOfset = t.HeightOfset - 1
	case "j":
		t.WidhtOfset = t.WidhtOfset + 1
	case "k":
		t.WidhtOfset = t.WidhtOfset - 1
	}

}
