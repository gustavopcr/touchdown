package internal

import (
	"fmt"
	"strings"
)

var count = make([]int, 20)
var points = [...]int{3, 6, 7, 8}

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
	count[0] = 1
	countPos()
}

func GetPos(n int) int {
	if(n>len(count)){
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
