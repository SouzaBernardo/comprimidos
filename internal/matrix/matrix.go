package matrix

import (
	"fmt"
	"image/png"
	"os"
)

type Pixel []uint8 // [R, G, B]

type Matrix struct {
	Pixels [][]Pixel
}

func NewMatrix(source string) *Matrix {

	file, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	pixels := make([][]Pixel, height)
	for y := 0; y < height; y++ {
		pixels[y] = make([]Pixel, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pixels[y][x] = Pixel{
				uint8(r >> 8),
				uint8(g >> 8),
				uint8(b >> 8),
			}
		}
	}
	return &Matrix{pixels}
}

func (m *Matrix) ForEachPixel(callback func(p *Pixel)) {
	for y := 0; y < len(m.Pixels); y++ {
		for x := 0; x < len(m.Pixels[y]); x++ {
			p := &(m.Pixels[y][x])
			callback(p)
		}
	}
}

func (m *Matrix) ForEachPixelIndex(callback func(y, x int, p *Pixel)) {
	for y := 0; y < len(m.Pixels); y++ {
		for x := 0; x < len(m.Pixels[y]); x++ {
			p := &(m.Pixels[y][x])
			callback(y, x, p)
		}
	}
}

func (m *Matrix) Print() {
	m.ForEachPixel(func(p *Pixel) {
		fmt.Println("Pixel:", *p)
	})
}
