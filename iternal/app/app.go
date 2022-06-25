package app

import (
	"bufio"
	"image"
	"os"

	"github.com/ybrbnf2080/genRand/iternal/convert"
	"github.com/ybrbnf2080/genRand/iternal/transform"
)

func init() {

}

func App(image image.Image, transform transform.Transform, draw func([][]int)) {
	pictDefoult := convert.Convert(image)
	//dotPict:=ViewNormal(pict, 60000)
	//dotPict2:= dotViewer(dotPict)
	pict := convert.Wider(pictDefoult)
	pict = convert.Compress(pict, transform.CompressCoof)
	pict = convert.Crop(pict, &transform)
	draw(pict)
	stdin := *bufio.NewReaderSize(os.Stdin, 64)
	for {
		transform.Controler(stdin)
		go func() {
			pict := convert.Wider(pictDefoult)
			pict = convert.Compress(pict, transform.CompressCoof)
			pict = convert.Crop(pict, &transform)
			draw(pict)
		}()
	}

}
