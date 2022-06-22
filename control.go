package main

import (
	"bufio"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	//	"sort"
)

type control struct {
	CompressCoof int
}

func controler(stdin bufio.Reader) {

	b := nextByte(&stdin)
	//fmt.Println("I got the byte", b, "("+string(b)+")")
	switch string(b) {

	case "+":
		CompressCoof = CompressCoof - 1
	case "-":
		CompressCoof = CompressCoof + 1
	case "q":
		os.Exit(1)
	case "l":
		HeightOfset = HeightOfset + 1
	case "h":
		HeightOfset = HeightOfset - 1
	case "j":
		WidhtOfset = WidhtOfset + 1
	case "k":
		WidhtOfset = WidhtOfset - 1
	}

}
