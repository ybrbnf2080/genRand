package render

import (
	"sort"
)

func sybmolSel(colorMap map[int]string, ColorShift int) func(int) string {
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
