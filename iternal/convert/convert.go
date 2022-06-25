package convert

import (
	"github.com/ybrbnf2080/genRand/iternal/transform"
)

func Crop(pict [][]int, transform *transform.Transform) [][]int {
	// Calc croped area
	endHeight := (len(pict[0]) - transform.Height) / 2
	endWidth := (len(pict) - transform.Width) / 2
	// check even`t size
	if (len(pict[0])-transform.Height)%2 != 0 {
		endHeight = endHeight + 1
	}
	if (len(pict[0])-transform.Width)%2 != 0 {
		endWidth = endWidth + 1
	}
	if endWidth-transform.WidhtOfset <= 0 {
		transform.WidhtOfset = transform.WidhtOfset - 1
	}
	if endHeight+transform.HeightOfset <= 0 {
		transform.HeightOfset = transform.HeightOfset + 1
	}
	if endWidth-transform.WidhtOfset <= 0 {
		transform.WidhtOfset = endWidth
	}
	if endHeight+transform.HeightOfset <= 0 {
		transform.HeightOfset = endHeight
	}
	// Crop
	if endWidth > 0 && transform.Width > 0 {
		if len(pict) < len(pict)-endWidth-transform.WidhtOfset {

		} else {
			pict = pict[endWidth-transform.WidhtOfset : len(pict)-endWidth-transform.WidhtOfset]
		}

	}

	if endHeight > 0 && transform.Height > 0 {
		for i, row := range pict {
			if len(row) <= 0 {
				pict[i] = make([]int, len(pict[0]))
				continue
			}
			if len(row) < len(row)-endHeight+transform.HeightOfset {
				//row = append(pict[i], 0)
			}
			pict[i] = row[endHeight+transform.HeightOfset : len(row)-endHeight+transform.HeightOfset]
		}
	}
	return pict
}

func Wider(pict [][]int) [][]int {
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
