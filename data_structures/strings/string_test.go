package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagramIndices(t *testing.T) {
	w := "ab"
	s := "abxaba"
	got := AnagramIndices(w, s)
	assert.Equal(t, []int{0, 3, 4}, got)
}

func TestIsPalindromePairs(t *testing.T) {
	input := []string{"code", "edoc", "da", "d"}
	got := IsPalindromePairs(input)
	assert.Equal(t, [][]int{{0, 1}, {1, 0}, {2, 3}}, got)
}

func TestHashUniqueUnicodeSort(t *testing.T) {
	input := "abcd"
	got := HashUniqueUnicodeSort(input)
	assert.Equal(t, true, got)
}

func TestHasUniqueUnicodeHash(t *testing.T) {
	input := "abcd"
	got := HasUniqueUnicodeHash(input)
	assert.Equal(t, true, got)
}

func TestHashUniqueASCIIChars(t *testing.T) {
	input := "abcd"
	got := HashUniqueASCIIChars(input)
	assert.Equal(t, true, got)
}
