package three_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkMergeSortedSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeSortedSlices([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	}
}

func TestMergeSortedSlices(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	mergeSortedSlices(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
}
