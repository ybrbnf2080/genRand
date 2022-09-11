package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	"github.com/ybrbnf2080/genRand/config"
	"github.com/ybrbnf2080/genRand/iternal/app"
	"github.com/ybrbnf2080/genRand/iternal/render"
)

func main() {
	var Transforms = config.Transforms
	var drawer = render.NewDrawer(render.SybmolSel(config.ColorMap, config.ColorShift))
	//ColorShift = 000
	imagePath := flag.Arg(0) //| "boobs.jpg" //"lol.jpg"l

	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.Decode(file)
	file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
		os.Exit(1)
	}

	if Transforms.CompressCoof == -1 && Transforms.Height > 0 {
		Transforms.CompressCoof = (image.Bounds().Dx() / Transforms.Height) * 2
	}
	app.App(image, Transforms, drawer)
	fmt.Println("Width:", image.Bounds().Dx(), "Height:", image.Bounds().Dy())
	fmt.Println("CompressCoof:", Transforms.CompressCoof, "ColorShift:", config.ColorShift)

}
