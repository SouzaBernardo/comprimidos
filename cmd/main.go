package main

import (
	"fmt"

	"github.com/SouzaBernardo/dip/internal/matrix"
	"github.com/SouzaBernardo/dip/internal/preprocessing"
	"github.com/SouzaBernardo/dip/internal/processing"
	"github.com/SouzaBernardo/dip/pkg/image"
)

const path = "assets/Comprimidos_%d.png"
const testImagePath = "test/image_%d.png"

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
		imagePath := fmt.Sprintf(path, i+1)
		fmt.Println("\nProcessando imagem:", imagePath)

		m := matrix.NewMatrix(imagePath)
		preprocessing.Exec(m)
		if err := image.Save(m, fmt.Sprintf(testImagePath, i+1)); err != nil {
			fmt.Println(err)
			return
		}
		total, broken, capsules, rounds := processing.Exec(m)

		fmt.Printf("\n================ Comparação dos Valores ================\n")
		fmt.Printf("Total     → Esperado: %3d | Recebido: %3d | Iguais: %t\n", total, value.total, total == value.total)
		fmt.Printf("Quebrados → Esperado: %3d | Recebido: %3d | Iguais: %t\n", broken, value.broken, broken == value.broken)
		fmt.Printf("Cápsulas  → Esperado: %3d | Recebido: %3d | Iguais: %t\n", capsules, value.capsules, capsules == value.capsules)
		fmt.Printf("Redondos  → Esperado: %3d | Recebido: %3d | Iguais: %t\n", rounds, value.rounds, rounds == value.rounds)
		fmt.Println("========================================================")

	}
}
