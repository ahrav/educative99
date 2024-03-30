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

func BenchmarkProductOfOtherElements(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = ProductOfOtherElements(input)
	}
}

func BenchmarkProductOfOtherElementsNoDivide(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		_ = ProductOfOtherElementsNoDivide(input)
	}
}

func TestSmallestWindowSorted(t *testing.T) {
	input := []int{3, 7, 5, 6, 9}
	begin, end := SmallestWindowSorted(input)
	assert.Equal(t, 1, begin)
	assert.Equal(t, 3, end)

}

func TestMaxSubarraySum(t *testing.T) {
	input := []int{24, -50, 42, 14, -5, 86}
	got := MaxSubarraySum(input)
	assert.Equal(t, 137, got)
}

func TestMaxSubarrayCircularSum(t *testing.T) {
	input := []int{8, -1, 3, 4}
	got := MaxSubarrayCircularSum(input)
	assert.Equal(t, 15, got)
}
