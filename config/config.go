package config

import (
	"flag"
	"fmt"
  "os"
	"os/exec"
	"strconv"
	"strings"

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
var Transforms transform.Transform
var ColorShift int
var Static bool

func init() {
	var Widht, Height, CompressCoof int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	size := string(out)
	if err != nil {
		fmt.Print("Error parse size you terminal")
		Height = 40
		Widht = 40
		
	} else {
		fmt.Print("l")
		slice := strings.Fields(size)
		Widht, _ = strconv.Atoi(slice[0])
		Height, _ = strconv.Atoi(slice[1])
	}
	//Widht, Height = 20, 50

	flag.IntVar(&CompressCoof, "comp", -1, "compression cooficent")
	flag.IntVar(&ColorShift, "color", 0, "color shift(exsposition)")

	flag.IntVar(&Height, "h", Height, "height crop resolution")
	flag.IntVar(&Widht, "w", Widht, "Widht crop resolution")
	flag.BoolVar(&Static, "s", false, "Static image rerturn")
	flag.Parse()
	Transforms = transform.NewTransform(Height, Widht, CompressCoof)

}
