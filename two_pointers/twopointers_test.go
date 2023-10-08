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

func TestRemoveNthLastNode(t *testing.T) {
	testCases := []struct {
		listInfo func() *ListNode
		n        int
		want     *LinkedList
	}{
		{
			listInfo: func() *ListNode {
				ll := InitLinkedList([]int{1, 2, 3, 4, 5})
				return ll.Head
			},
			n:    2,
			want: InitLinkedList([]int{1, 2, 3, 5}),
		},
		{
			listInfo: func() *ListNode {
				ll := InitLinkedList([]int{1})
				return ll.Head
			},
			n:    1,
			want: nil,
		},
		{
			listInfo: func() *ListNode {
				ll := InitLinkedList([]int{1, 2})
				return ll.Head
			},
			n:    1,
			want: InitLinkedList([]int{1}),
		},
		{
			listInfo: func() *ListNode {
				ll := InitLinkedList([]int{1, 2})
				return ll.Head
			},
			n:    2,
			want: InitLinkedList([]int{2}),
		},
		{
			listInfo: func() *ListNode {
				return nil
			},
			n:    1,
			want: nil,
		},
	}

	for _, tc := range testCases {
		head := tc.listInfo()
		got := removeNthLastNode(head, tc.n)
		if tc.want == nil {
			assert.Nil(t, got)
			continue
		}
		assert.Equal(t, tc.want.Display(), ConstructLinkedListFromNode(got).Display())
	}
}
