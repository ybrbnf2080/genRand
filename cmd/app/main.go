package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/ybrbnf2080/genRand/config"
	"github.com/ybrbnf2080/genRand/iternal/app"
	"github.com/ybrbnf2080/genRand/iternal/render"
	"github.com/ybrbnf2080/genRand/iternal/transform"
)

var Transforms transform.Transform
var ColorShift int

func init() {
	var Widht, Height, CompressCoof int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	size := string(out)
	if err != nil {
		fmt.Print("Error parse size you terminal")
		Height = 40
		return
	}
	slice := strings.Fields(size)
	Widht, _ = strconv.Atoi(slice[0])
	Height, _ = strconv.Atoi(slice[1])
	//Widht, Height = 20, 50

	flag.IntVar(&CompressCoof, "comp", -1, "compression cooficent")
	flag.IntVar(&ColorShift, "color", 0, "color shift(exsposition)")

	flag.IntVar(&Height, "h", Height, "height crop resolution")
	flag.IntVar(&Widht, "w", Widht, "Widht crop resolution")
	flag.Parse()
	Transforms = transform.NewTransform(Height, Widht, CompressCoof)
	         
}

func main() {
	var drawer = render.NewDrawer(render.SybmolSel(config.ColorMap, ColorShift))
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
	fmt.Println("CompressCoof:", Transforms.CompressCoof, "ColorShift:", ColorShift)

}
