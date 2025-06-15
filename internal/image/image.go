package image

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/SouzaBernardo/dip/internal/matrix"
)


func Save(m *matrix.Matrix) error {
	height := len(m.Pixels)
	if height == 0 {
		return nil // ou erro, se preferir
	}
	width := len(m.Pixels[0])

	// Cria uma nova imagem RGBA
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Preenche os pixels
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := m.Pixels[y][x]
			if len(pixel) < 3 {
				pixel = []uint8{0, 0, 0} // fallback pra evitar panic
			}
			r, g, b := pixel[0], pixel[1], pixel[2]
			img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}

	// Cria o arquivo
	file, err := os.Create("./test/image.png")
	if err != nil {
		return err
	}
	defer file.Close()

	// Codifica e salva como PNG
	return png.Encode(file, img)
}