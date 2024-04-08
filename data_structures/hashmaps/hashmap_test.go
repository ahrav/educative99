package hashmaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	input1 := []int{137, 59, 92, 122, 52, 131, 79, 236}
	input2 := []int{137, 59, 92, 122, 52, 131, 79, 236}
	want := []int{236, 92, 122, 131, 131, 236, 236, -1}
	got := nextGreaterElement(input1, input2)
	assert.Equal(t, want, got)
}

func TestFractionToDecimal(t *testing.T) {
	num, denom := 8, 666
	got := fractionToDecimal(num, denom)
	assert.Equal(t, "0.(012)", got)
}
