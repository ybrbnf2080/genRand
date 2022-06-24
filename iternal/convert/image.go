package convert

import (
	"bufio"
	"image"
	"os"
)

func convert(image image.Image) [][]int {
	resolution := image.Bounds()
	var pict [][]int
	for i := 1; i < resolution.Dy(); i++ {
		var row []int
		for w := 1; w < resolution.Dx(); w++ {
			pix := image.At(w, i)
			r, g, b, _ := pix.RGBA()
			summ := (r + g + b)
			row = append(row, int(summ))
		}
		pict = append(pict, row)

	}
	return pict
}

func draw(pict [][]int, sybmolisator func(int) string) {
	state := make([]string, len(pict))
	//start := time.Now()

	printer := *bufio.NewWriter(os.Stdout)

	for i := range state {
		go func(i int) {
			var row string
			for _, w := range pict[i] {
				summ := w
				//if "" == unic[summ] {
				//	unic[summ] = string(rune(len(unic)))
				//}
				row = row + sybmolisator(summ)

			}
			state[i] = row
		}(i)
	}

	for i := range state {
		printer.WriteString(state[i] + "\n")
	}
	printer.Flush()
	//fmt.Println("\n", time.Since(start))

}
