package convert

import (
	"github.com/ybrbnf/genRand/iternal/transform"
)

func crop(pict [][]int, transform Transform.Transform) [][]int {
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
