package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por realizar a média aritmética dos valores passados para a função
func Media(valores ...float64) float64 {
	total := 0.0

	for i := 0; i < len(valores); i++ {
		total += valores[i]
	}

	media := total / float64(len(valores))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)

	return mediaArredondada
}
