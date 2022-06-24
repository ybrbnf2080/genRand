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
)



func drawHorisontal(pict [][]int, sybmolisator func(int) string, olorPallet map[int]string) {

	for i, _ := range pict[0] {
		for _, w := range pict {
			//fmt.Print(i)
			summ := w[i]
			//if "" == unic[summ] {
			//	unic[summ] = string(rune(len(unic)))
			//}

			fmt.Print(sybmolisator(summ))
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
	// Calc croped area
	endHeight := (len(pict[0]) - height) / 2
	endWidth := (len(pict) - widht) / 2
	// check even`t size
	if (len(pict[0])-height)%2 != 0 {
		endHeight = endHeight + 1
	}
	if (len(pict[0])-widht)%2 != 0 {
		endWidth = endWidth + 1
	}
	if endWidth-WidhtOfset <= 0 {
		WidhtOfset = WidhtOfset - 1
	}
	if endHeight+HeightOfset <= 0 {
		HeightOfset = HeightOfset + 1
	}
	if endWidth-WidhtOfset <= 0 {
		WidhtOfset = endWidth
	}
	if endHeight+HeightOfset <= 0 {
		HeightOfset = endHeight
	}
	// Crop
	if endWidth > 0 && widht > 0 {
		if len(pict) < len(pict)-endWidth-WidhtOfset {

		} else {
			pict = pict[endWidth-WidhtOfset : len(pict)-endWidth-WidhtOfset]
		}

	}

	if endHeight > 0 && height > 0 {
		for i, row := range pict {
			if len(row) <= 0 {
				pict[i] = make([]int, len(pict[0]))
				continue
			}
			if len(row) < len(row)-endHeight+HeightOfset {
				//row = append(pict[i], 0)
			}
			pict[i] = row[endHeight+HeightOfset : len(row)-endHeight+HeightOfset]
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
