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
