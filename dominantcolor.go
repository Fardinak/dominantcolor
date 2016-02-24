package dominantcolor

import (
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/color/palette"
)

// Extract returns the dominant color of a given image
func Extract(img image.Image) color.Color {
	// Resize the image for faster computation
	image = resize.Resize(100, 0, image, resize.Bilinear)
	bounds := img.Bounds()

	colorCount := make([]int, len(palette.WebSafe))
	biggest := []int{0, 0}

	for i := 0; i <= bounds.Max.X; i++ {
		for j := 0; j <= bounds.Max.Y; j++ {
			pixel := img.At(i, j)
			c := color.Palette(palette.WebSafe).Index(pixel)
			colorCount[c]++

			// Find the most used color
			if colorCount[c] > biggest[1] {
				biggest[0] = c
				biggest[1] = colorCount[c]
			}
		}
	}

	return palette.WebSafe[biggest[0]]
}
