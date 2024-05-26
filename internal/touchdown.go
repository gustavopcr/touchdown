package internal

import (
	"fmt"
	"strings"
)

var count = make([]int, 20)
var points = [...]int{3, 6, 7, 8}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

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
	return count[n]
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
