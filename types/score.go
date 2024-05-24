package types

import (
	"fmt"
	"strings"
)

type Score struct {
	Score string `json:"score"`
}

func (s *Score) ParseScore() ([]int, error) {
	input := strings.TrimSpace(s.Score)
	input = strings.ToLower(input)
	var n1, n2 int
	_, err := fmt.Sscanf(input, "%dx%d", &n1, &n2)
	if err != nil {
		return nil, fmt.Errorf("error parsing input: %w", err)
	}

	return []int{n1, n2}, nil
}
