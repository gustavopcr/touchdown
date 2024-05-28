package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorParseInput(t *testing.T) {
	input := "1+15"
	_, err := ParseScore(input)

	assert.Error(t, err, "Parse string in incorrect format")
}

func TestParseInput(t *testing.T) {
	input := "3x15"
	want, err := ParseScore(input)

	assert.NoError(t, err)
	assert.Equal(t, 3, want[0])
	assert.Equal(t, 15, want[1])
}

func TestValuesGenerated(t *testing.T) {
	StartTouchdown()
	want := GetPos(15)

	assert.Equal(t, 4, want)
}
