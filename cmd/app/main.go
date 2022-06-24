package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)


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
