package linkedlists

import (
	"fmt"
	"math"
)

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = nil
	return node
}

type EduLinkedList struct {
	head *EduLinkedListNode
}

// InsertNodeAtHead method will insert a LinkedListNode at head of a linked list.
func (l *EduLinkedList) InsertNodeAtHead(node *EduLinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

// CreateLinkedList method will create the linked list using the given integer array
// with the help of InsertAthead method.
func (l *EduLinkedList) CreateLinkedList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		newNode := InitLinkedListNode(lst[i])
		l.InsertNodeAtHead(newNode)
	}
}

// DisplayLinkedList method will display the elements of linked list.
func (l *EduLinkedList) DisplayLinkedList() {
	temp := l.head
	fmt.Print("[")
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

func reverse(head *EduLinkedListNode) *EduLinkedListNode {
	if head == nil || head.next == nil {
		return head
	}

	var prev, curr, next *EduLinkedListNode = nil, head, nil
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}

func reverseEvenLengthGroups(head *EduLinkedListNode) {
	// Write your code here
}

func reorderList(head *EduLinkedListNode) {
	if head == nil || head.next == nil {
		return
	}
	one, two := split(head)
	two = reverse(two)
	mergeAlternating(one, two)
}

// merge two linked lists in alternating order starting with the first list.
// The lists will either be the same size or have one extra node in the second list.
// If the second list has an extra node, it should be merged to the end of the merged list.
func mergeAlternating(head1 *EduLinkedListNode, head2 *EduLinkedListNode) *EduLinkedListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	one, two := head1, head2
	var temp *EduLinkedListNode
	for one != nil && two != nil {
		temp = one.next
		one.next = two
		one = temp

		temp = two.next
		two.next = one
		two = temp
	}
	// if one list is longer than the other we simply append.
	// Handle any extra nodes
	tail := one
	if two != nil {
		tail = two
	}

	// Find the last node in the merged list
	temp = head1
	for temp.next != nil {
		temp = temp.next
	}

	// If there's an extra node, append it
	if tail != nil {
		temp.next = tail
	}
	return head1
}

// split a list into two lists. Keep the first half in the original list
// the extra element should be in the second list.
func split(head *EduLinkedListNode) (*EduLinkedListNode, *EduLinkedListNode) {
	if head == nil {
		return nil, nil
	}

	slow, fast := head, head
	var prev *EduLinkedListNode
	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	prev.next = nil
	return head, slow
}

func getMiddleNode(head *EduLinkedListNode) *EduLinkedListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	slow.next = nil

	return slow
}

func rotateLeft(d int32, arr []int32) []int32 {
	if len(arr) == 0 || int32(len(arr)) == d {
		return arr
	}

	d = d % int32(len(arr))
	res := make([]int32, 0, len(arr))
	res = append(res, arr[d:]...)
	res = append(res, arr[:d]...)

	return res
}

func rotateLeftInPlace(d int32, arr []int32) []int32 {
	if len(arr) == 0 || int32(len(arr)) == d {
		return arr
	}

	d = d % int32(len(arr))
	reverseArrayInPlace(arr[:d])
	reverseArrayInPlace(arr[d:])
	reverseArrayInPlace(arr)

	return arr
}

func reverseArrayInPlace(arr []int32) []int32 {
	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	return arr
}

func powerSum(x int32, n int32) int32 {
	return powerSumHelper(x, n, 1)
}

func powerSumHelper(x int32, n int32, num int32) int32 {
	p := int32(math.Pow(float64(num), float64(n)))
	if p > x {
		return 0
	} else if p == x {
		return 1
	} else {
		return powerSumHelper(x, n, num+1) + powerSumHelper(x-p, n, num+1)
	}
}

var memo = make(map[string]int32)

func powerSumHelperMemo(x, n, num int32) int32 {
	key := fmt.Sprintf("%d:%d:%d", x, n, num)
	if val, exists := memo[key]; exists {
		return val
	}

	p := int32(math.Pow(float64(num), float64(n)))

	if p > x {
		return 0
	} else if p == x {
		return 1
	} else {
		memo[key] = powerSumHelperMemo(x, n, num+1) + powerSumHelperMemo(x-p, n, num+1)
		return memo[key]
	}
}

func powerSumDP(x, n int32) int32 {
	// Initialize DP array with size x+1. dp[i] will store the number of ways to make the sum i.
	dp := make([]int32, x+1)
	dp[0] = 1 // Base case: one way to make the sum 0: use no numbers at all.

	// Loop through each 'num' to see if it should be included in the sum.
	for num := int32(1); int32(math.Pow(float64(num), float64(n))) <= x; num++ {
		p := int32(math.Pow(float64(num), float64(n)))

		// Update dp array: try including 'num'
		for i := x; i >= p; i-- {
			dp[i] += dp[i-p]
		}
	}

	return dp[x] // This contains the number of ways to make the sum x.
}

// canConstruct returns true if searchWord can be constructed from the letters in corpus.
func canConstruct(searchWord, corpus string) bool {
	if len(searchWord) == 0 {
		return true
	}

	if len(corpus) == 0 {
		return false
	}

	if len(searchWord) > len(corpus) {
		return false
	}

	charCount := make(map[rune]int)
	remainingChars := len(searchWord)

	for _, c := range searchWord {
		charCount[c]++
	}

	for _, c := range corpus {
		if count, exists := charCount[c]; exists && count > 0 {
			charCount[c]--
			remainingChars--

			if charCount[c] == 0 {
				delete(charCount, c)
			}

			if remainingChars == 0 {
				return true
			}
		}
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ConvertBSTToLinkedList traverses the tree in-order and appends nodes to the list.
func ConvertBSTToLinkedList(root *TreeNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	inorder(root, &prev)
	return dummy.Next
}

// inorder recursively performs an in-order traversal.
func inorder(root *TreeNode, prev **ListNode) {
	if root == nil {
		return
	}
	inorder(root.Left, prev)
	(*prev).Next = &ListNode{Val: root.Val}
	*prev = (*prev).Next
	inorder(root.Right, prev)
}
