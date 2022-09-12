package convert

import (
	"image"
	"image/jpeg"
)

func init() {
	// damn important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

}
func Convert(image image.Image) [][]int {
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
