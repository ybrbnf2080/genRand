package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sync"
	"time"
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
	var wg sync.WaitGroup
	wg.Add(len(pict))
	start := time.Now()

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
			wg.Done()
		}(i)
	}

	wg.Wait()
	for i := range state {
		printer.WriteString(state[i] + "\n")
	}
	printer.Flush()
	fmt.Println("\n", time.Since(start))

}

func drawHorisontal(pict [][]int, sybmolisator func(int) string, olorPallet map[int]string) {

	for i, _ := range pict[0] {
		for _, w := range pict {
			//fmt.Print(i)
			summ := w[i]
			//if "" == unic[summ] {
			//	unic[summ] = string(rune(len(unic)))
			//}

			fmt.Print( sybmolisator(summ, ))
			fmt.Print(sybmolisator(summ))
		}
		fmt.Println("")

	}

}

func compress(pict [][]int, compCoof int) [][]int {
	// endWidth := len(pict) / 2
	//endHeight := len(pict[0]) / 2
	if compCoof <= 1 {
		return pict
	}
	var endPic [][]int
	for i := 0; i <= len(pict)-compCoof; i = i + compCoof {
		var endRow []int
		for l := 0; l <= len(pict[0])-compCoof; l = l + compCoof {
			color := func() int {

				endColor := 0
				//(pict[i][l] + pict[i+1][l] + pict[i][l+1] + pict[i+1][l+1]) / 4

				for i2 := 0; i2 <= compCoof-1; i2++ {
					for l2 := 0; l2 <= compCoof-1; l2++ {
						endColor = endColor + pict[i+i2][l+l2]

					}
				}

				return endColor / (compCoof * compCoof)
			}()

			endRow = append(endRow, color)
		}
		endPic = append(endPic, endRow)
	}
	return endPic
}

func crop(pict [][]int, height int, widht int) [][]int {
	endHeight := (len(pict[0]) - height) / 2
	endWidth := (len(pict) - widht) / 2
	if (len(pict[0])-height)%2 != 0 {
		endHeight = endHeight + 1
	}
	if (len(pict[0])-widht)%2 != 0 {
		endWidth = endWidth + 1
	}
	if endWidth > 0 && widht > 0 {
		pict = pict[endWidth : len(pict)-endWidth]
	}

<<<<<<< HEAD
=======
	if endHeight > 0 && height > 0{
		endWidth = endWidth + 1
	}
	if endWidth > 0 && widht > 0 {
		pict = pict[endWidth : len(pict)-endWidth]
	}

>>>>>>> 64d12ee (lol)
	if endHeight > 0 && height > 0 {

		for i, row := range pict {
			pict[i] = row[endHeight : len(row)-endHeight]
		}
	}
	return pict
}

func wider(pict [][]int) [][]int {
	var widePict [][]int
	for _, h := range pict {
		var row []int
		for _, w := range h {
			summ := w
			//if "" == unic[summ] {
			//	unic[summ] = string(rune(len(unic)))
			//}
			row = append(row, summ, summ)

		}
		widePict = append(widePict, row)
	}
	return widePict
}
