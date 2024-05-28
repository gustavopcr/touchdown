package internal

import (
	"fmt"
	"strings"
)

var count = make([]int, 20)
var points = [...]int{3, 6, 7, 8}

/*
utiliza de programacao dinamica(problema da moeda) para contar todas as possibilidades. A ideia é construir a solucao moeda por moeda de forma a aproveitar
computacoes anteriores. Exemplo: utilizando a moeda 3, existe 1 solução para 3,6 e 9. Utilizando a moeda 2, 6 e 9 ganham mais uma solução
*/
func countPos() {
	for _, p := range points {
		for i := 1; i < len(count); i++ {
			if i-p >= 0 {
				count[i] += count[i-p]
			}
		}
	}
}

func StartTouchdown() {
	count[0] = 1 // necessario indice 0 ser 1 pois existem apenas uma maneira de alcancar 0, que é escolhendo nenhuma nao pontuar
	countPos()
}

func GetPos(n int) int {
	if n > len(count) { // se numero for maior que slice atual, aumenta tamanho do slice
		count = make([]int, n+1)
		StartTouchdown()
	}
	return count[n]
}

func GetAll() []int {
	return count
}
func ParseScore(s string) ([]int, error) {
	input := strings.TrimSpace(s)
	input = strings.ToLower(input)
	var n1, n2 int
	_, err := fmt.Sscanf(input, "%dx%d", &n1, &n2)
	if err != nil {
		return nil, fmt.Errorf("error parsing input: %w", err)
	}

	return []int{n1, n2}, nil
}
