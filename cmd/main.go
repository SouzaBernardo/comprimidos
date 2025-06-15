package main

import (
	"fmt"

	"github.com/SouzaBernardo/dip/internal/matrix"
	"github.com/SouzaBernardo/dip/internal/preprocessing"
	"github.com/SouzaBernardo/dip/internal/processing"
)

const path = "assets/Comprimidos_%d.png"

/*
1. O total de comprimidos na esteira.
2. Quantos comprimidos estão quebrados e não podem ser aproveitados.
3. Quantas são as cápsulas e quantos são os comprimidos redondos.
*/

type resultExpect struct {
	total    int // O total de comprimidos na esteira.
	broken   int // Quantos comprimidos estão quebrados e não podem ser aproveitados.
	capsules int // Quantas são as cápsulas
	rounds   int // Quantos são os comprimidos redondos
}

var results [5]resultExpect = [5]resultExpect{
	{14, 0, 3, 11},
	{12, 0, 3, 9},
	{10, 2, 2, 6},
	{8, 1, 2, 5},
	{11, 0, 3, 8},
}

func main() {
	for i, value := range results {

		imagePath := fmt.Sprintf(path, i + 1)
		m := matrix.NewMatrix(imagePath)

		preprocessing.Exec(m)
		w, x, y, z := processing.Exec(m)

		fmt.Println("Image:", imagePath)
		fmt.Println(w == value.total)
		fmt.Println(x == value.broken)
		fmt.Println(y == value.capsules)
		fmt.Println(z == value.rounds)
	}
}
