package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductOfOtherElements(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := []int{120, 60, 40, 30, 24}

	got := ProductOfOtherElements(input)
	got2 := ProductOfOtherElementsNoDivide(input)
	assert.Equal(t, want, got)
	assert.Equal(t, want, got2)
}
