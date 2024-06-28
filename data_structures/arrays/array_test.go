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

func TestSmallerCounts(t *testing.T) {
	input := []int{3, 4, 9, 6, 1}
	want := []int{1, 1, 2, 1, 0}
	got := SmallerCounts(input)
	assert.Equal(t, want, got)
}

func TestMinRemovals(t *testing.T) {
	input := []int{6, 9, 6, 7, 2, 7, 2}
	k := 2
	th := 13
	got := minRemovals(input, k, th)
	assert.Equal(t, 2, got)
}

func TestRemoveElement(t *testing.T) {
	input := []int{3, 2, 2, 3}
	val := 3
	want := 2
	got := removeElement(input, val)
	assert.Equal(t, want, got)
}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1, 2}
	want := 2
	got := removeDuplicates(input)
	assert.Equal(t, want, got)
	assert.Equal(t, []int{1, 2}, input)
}

func TestRemoveDuplicatesTwo(t *testing.T) {
	input := []int{1, 1, 1, 2, 2, 2, 3}
	want := 5
	got := removeDuplicatesTwo(input)
	assert.Equal(t, want, got)
	assert.Equal(t, []int{1, 1, 2, 2, 3}, input[:got])
}
