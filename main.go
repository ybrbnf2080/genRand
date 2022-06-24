package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"os/exec"

	//	"sort"
	"strconv"
	"strings"
)

func init() {
	// damn important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

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
var ColorShift = 16000
var CompressCoof = 0

func forInImage(image image.Image) {
	pictDefoult := convert(image)
	//dotPict:=ViewNormal(pict, 60000)
	//dotPict2:= dotViewer(dotPict)
	pict := wider(pictDefoult)
	pict = compress(pict, CompressCoof)
	pict = crop(pict, Height, Widht)
	draw(pict, sybmolSel(ColorMap))
	stdin := *bufio.NewReaderSize(os.Stdin, 64)
	for {
		controler(stdin)
		go func() {
			pict = wider(pictDefoult)
			pict = compress(pict, CompressCoof)
			pict = crop(pict, Height, Widht)
			draw(pict, sybmolSel(ColorMap, ColorShift))
		}()
	}

}

var Widht, Height, HeightOfset, WidhtOfset int

func init() {
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
}
func init() {

	flag.IntVar(&CompressCoof, "comp", -1, "compression cooficent")
	flag.IntVar(&ColorShift, "color", 0, "color shift(exsposition)")

	flag.IntVar(&Height, "h", Height, "height crop resolution")
	flag.IntVar(&Widht, "w", Widht, "Widht crop resolution")

	flag.Parse()
}

func main() {
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
	}

	if CompressCoof == -1 && Height > 0 {
		CompressCoof = (image.Bounds().Dx() / Height) * 2
	}

	forInImage(image)
	fmt.Println("Width:", image.Bounds().Dx(), "Height:", image.Bounds().Dy())
	fmt.Println("CompressCoof:", CompressCoof, "ColorShift:", ColorShift)

}
