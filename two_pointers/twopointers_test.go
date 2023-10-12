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

func TestSortColors(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{2, 0, 2, 1, 1, 0},
			want:  []int{0, 0, 1, 1, 2, 2},
		},
		{
			input: []int{2, 0, 1},
			want:  []int{0, 1, 2},
		},
		{
			input: []int{0},
			want:  []int{0},
		},
		{
			input: []int{1, 1, 1},
			want:  []int{1, 1, 1},
		},
		{
			input: nil,
			want:  nil,
		},
	}

	for _, tc := range testCases {
		sortColors(tc.input)
		assert.Equal(t, tc.want, tc.input)
	}
}

func TestIsHappy(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{input: 19, want: true},
		{input: 2, want: false},
		{input: 4, want: false},
		{input: 0, want: false},
		{input: 28, want: true},
		{input: 1, want: true},
	}

	for _, tc := range testCases {
		got := isHappy(tc.input)
		assert.Equal(t, tc.want, got)
	}
}

func BenchmarkIsHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isHappy(1912)
	}
}

func TestIsHappyTwoPointer(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{input: 19, want: true},
		{input: 2, want: false},
		{input: 4, want: false},
		{input: 0, want: false},
		{input: 28, want: true},
		{input: 1, want: true},
	}

	for _, tc := range testCases {
		got := isHappyTwoPointer(tc.input)
		assert.Equal(t, tc.want, got)
	}
}

func BenchmarkIsHappyTwoPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isHappyTwoPointer(1912)
	}
}

func TestCircularArrayLoop(t *testing.T) {
	testCases := []struct {
		input []int
		want  bool
	}{
		{
			input: []int{2, -1, 1, 2, 2},
			want:  true,
		},
		{
			input: []int{-1, 2},
			want:  false,
		},
		{
			input: []int{-2, 1, -1, -2, -2},
			want:  false,
		},
		{
			input: []int{-1, -2, -3, -4, -5},
			want:  false,
		},
		{
			input: []int{1, 1, 2},
			want:  true,
		},
		{
			input: []int{2, 2, 2, 2, 2, 4, 7},
			want:  false,
		},
		{
			input: []int{-1, -1, -1},
			want:  true,
		},
		{
			input: []int{-1, -1, -2},
			want:  true,
		},
	}

	for _, tc := range testCases {
		got := circularArrayLoop(tc.input)
		assert.Equal(t, tc.want, got)
	}
}

func BenchmarkCircularArrayLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = circularArrayLoop([]int{2, -1, 1, 2, 2})
	}
}

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		want   int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			target: 2,
			want:   1,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			target: 4,
			want:   3,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			target: 12,
			want:   -1,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			target: 0,
			want:   -1,
		},
	}

	for _, tc := range testCases {
		got := binarySearch(tc.input, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
