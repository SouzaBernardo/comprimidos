package processing

import "github.com/SouzaBernardo/dip/internal/matrix"

const CAPSULES = "capsule"
const ROUND = "round"
const BROKEN = "broken"
const CAPSULES_RATIO = 1.2
const ROUND_RATIO = 1.1

type Point struct {
	Y, X int
}

func Exec(m *matrix.Matrix) (int, int, int, int) {

	total := 0
	broken := 0
	capsules := 0
	rounds := 0

	visited := make([][]bool, len(m.Pixels))
	for i := range visited {
		visited[i] = make([]bool, len(m.Pixels[0]))
	}

	m.ForEachPixelIndex(func(y, x int, p *matrix.Pixel) {
		if isWhite(p) && !visited[y][x] {
			region := floodFill(m, y, x, visited)
			typeRegion := classifyRegion(region)
			if typeRegion == CAPSULES {
				capsules++
			} else if typeRegion == ROUND {
				rounds++
			} else {
				broken++
			}
			total++
		}
	})

	return total, broken, capsules, rounds
}

func isWhite(p *matrix.Pixel) bool {
	return (*p)[0] == 255 && (*p)[1] == 255 && (*p)[2] == 255
}

func floodFill(m *matrix.Matrix, startY, startX int, visited [][]bool) []Point {
	var region []Point
	queue := []Point{{startY, startX}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		y, x := current.Y, current.X
		if y < 0 || y >= len(m.Pixels) || x < 0 || x >= len(m.Pixels[0]) || visited[y][x] || !isWhite(&m.Pixels[y][x]) {
			continue
		}

		visited[y][x] = true
		region = append(region, current)

		queue = append(queue,
			Point{y - 1, x},
			Point{y + 1, x},
			Point{y, x - 1},
			Point{y, x + 1},
		)
	}

	return region
}

func classifyRegion(region []Point) string {
	if len(region) == 0 {
		return ""
	}

	minX, maxX := region[0].X, region[0].X
	minY, maxY := region[0].Y, region[0].Y

	for _, pt := range region {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}

	width := maxX - minX + 1
	height := maxY - minY + 1

	aspectRatio := float64(width) / float64(height)
	if aspectRatio < 1.0 {
		aspectRatio = 1.0 / aspectRatio
	}

	if aspectRatio <= ROUND_RATIO {
		return ROUND
	} else if aspectRatio <= CAPSULES_RATIO {
		return CAPSULES
	}

	return BROKEN
}
