package convert

func Compress(pict [][]int, compCoof int) [][]int {
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
