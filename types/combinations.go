package types

type Combination struct {
	Combination int `json:"combinations"`
}

func NewCombination(n int) Combination {
	return Combination{
		Combination: n,
	}
}
