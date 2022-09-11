package render

import (
	"bufio"
	"os"
	"sync"
)

func NewDrawer(sybmolisator func(int) string) func([][]int) {
	//start := time.Now()

	printer := *bufio.NewWriter(os.Stdout)
	return func(pict [][]int) {
		var wg sync.WaitGroup

		wg.Add(len(pict))
		state := make([]string, len(pict))

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
	}
}

func Drawer(sybmolisator func(int) string) func([][]int) string {
	return func(pict [][]int) string{
		var wg sync.WaitGroup

		wg.Add(len(pict))
		state := make([]string, len(pict))

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
		var endPict string
		for _, r := range state {
			endPict = endPict + r + "\n"
		}
		return endPict
	}
}
