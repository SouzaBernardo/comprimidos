package image

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/SouzaBernardo/dip/internal/matrix"
)

func Save(m *matrix.Matrix, path string) error {
	height := len(m.Pixels)
	if height == 0 {
		return nil // ou erro, se preferir
	}
	width := len(m.Pixels[0])

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := m.Pixels[y][x]
			if len(pixel) < 3 {
				pixel = []uint8{0, 0, 0}
			}
			r, g, b := pixel[0], pixel[1], pixel[2]
			img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
