package preprocessing

import (
	"github.com/SouzaBernardo/dip/internal/matrix"
)

func Exec(m *matrix.Matrix) {
	applyThreshold(m)
}

func applyThreshold(m *matrix.Matrix) {
	m.ForEachPixel(func(p *matrix.Pixel) {
		r, g, b := (*p)[0], (*p)[1], (*p)[2]
		if r <= 120 && g <= 120 && b <= 120 {
			*p = matrix.Pixel{0, 0, 0}
		} else {
			*p = matrix.Pixel{255, 255, 255}
		}
	})
}
