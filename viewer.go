package main

import (
	"sort"

	"strconv"
)

func sybmolSel(colorMap map[int]string) func(int) string {
	keys := make([]int, 0)
	for k, _ := range colorMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return func(color int) string {

		for _, key := range keys {
			col := colorMap[key]
			//		fmt.Println(key, color, col)
			if color < key+ColorShift {
				return col
			}
		}
		return colorMap[0]
	}
}

func ViewNormal(pict [][]int, exspose int) [][]bool {
	var endPict [][]bool
	for _, row := range pict {
		var endRow []bool
		for _, pix := range row {
			endRow = append(endRow, pix > exspose)

			if len(endRow)%4 == 0 {
				for i := 0; i < len(endRow); i++ {
					endRow = append(endRow, false)
				}
			}
		}
		endPict = append(endPict, endRow)
	}
	if len(endPict)%4 == 0 {

		for i := 0; i < len(endPict); i++ {
			endPict = append(endPict, endPict[0])
		}
	}

	return endPict
}

func dotViewer(pict [][]bool) [][]int {
	var endPict [][]int
	for iRow := 0; iRow < len(pict); iRow = iRow + 4 {
		var row []int
		for i := 0; i < len(pict[0]); i = i + 2 {
			char := ""
			for iRow2 := i; iRow2 < iRow+4; iRow2++ {
				for i2 := i; i2 < i+2; i2++ {
					if pict[iRow2][i2] {
						char = char + "1"
					} else {
						char = char + "0"
					}
				}
			}
			dot, _ := strconv.ParseInt(char, 2, 0)

			row = append(row, int(dot))
		}
		endPict = append(endPict, row)
	}
	return endPict
}

func dotSel([]bool) string {
	return " "
}
