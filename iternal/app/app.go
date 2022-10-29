package app

import (
	"bufio"
	"image"
	"os"

	"github.com/ybrbnf2080/genRand/iternal/convert"
	"github.com/ybrbnf2080/genRand/iternal/transform"
)

var ColorMap = map[int]string{
	0:      "#",
	5000:   " ",
	10000:  ".",
	20000:  ",",
	22500:  "-",
	25000:  "=",
	27500:  "!",
	30000:  "|",
	40000:  "i",
	59000:  "t",
	79000:  "n",
	80000:  "s",
	99000:  "z",
	120000: "a",
	140000: "h",
	160815: "O",
	190036: "Q",
	190037: "@",
}

func init() {

}

func App(image image.Image, transform transform.Transform, draw func([][]int), static bool) {
	pictDefoult := convert.Convert(image)
	//dotPict:=ViewNormal(pict, 60000)
	//dotPict2:= dotViewer(dotPict)
	pict := convert.Wider(pictDefoult)
	pict = convert.Compress(pict, transform.CompressCoof)
	pict = convert.Crop(pict, &transform)
	stdin := *bufio.NewReaderSize(os.Stdin, 64)
	if static {
			pict := convert.Wider(pictDefoult)
			pict = convert.Compress(pict, transform.CompressCoof)
			pict = convert.Crop(pict, &transform)
			draw(pict)
			return
	}
	for {
		go func() {
			pict := convert.Wider(pictDefoult)
			pict = convert.Compress(pict, transform.CompressCoof)
			pict = convert.Crop(pict, &transform)
			draw(pict)
		}()
		transform.Controler(stdin)
	}

}
func Convert(image image.Image, transform transform.Transform, draw func([][]int) string) string {
	pictDefoult := convert.Convert(image)
	//dotPict:=ViewNormal(pict, 60000)
	//dotPict2:= dotViewer(dotPict)
	pict := convert.Wider(pictDefoult)
	pict = convert.Compress(pict, transform.CompressCoof)
	pict = convert.Crop(pict, &transform)
	return draw(pict)
}
