package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "kayak",
			want:  true,
		},
		{
			input: "racecar",
			want:  true,
		},
		{
			"not a palindrome",
			false,
		},
		{
			"a",
			true,
		},
		{
			"aa",
			true,
		},
		{
			"",
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := isPalindrome(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
