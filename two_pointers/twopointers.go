package two_pointers

func isPalindrome(s string) bool {
	if s == "" {
		return false
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// ListNode is a node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// InitListNode initializes a linked list from a single integer.
func InitListNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

// LinkedList is a linked list.
type LinkedList struct {
	Head *ListNode
}

// InsertAtHead inserts a node at the head of a linked list.
func (ll *LinkedList) InsertAtHead(node *ListNode) {
	if ll == nil {
		return
	}

	node.Next = ll.Head
	ll.Head = node
}

// ConstructLinkedListFromNode constructs a linked list from a node.
func ConstructLinkedListFromNode(node *ListNode) *LinkedList {
	if node == nil {
		return nil
	}

	return &LinkedList{Head: node}
}

// InitLinkedList initializes a linked list from a slice of integers.
func InitLinkedList(vals []int) *LinkedList {
	if len(vals) == 0 {
		return nil
	}

	head := InitListNode(vals[0])
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = InitListNode(vals[i])
		curr = curr.Next
	}

	return &LinkedList{Head: head}
}

// Display displays a linked list.
func (ll *LinkedList) Display() []int {
	if ll == nil {
		return nil
	}

	var vals []int
	curr := ll.Head
	for curr != nil {
		vals = append(vals, curr.Val)
		curr = curr.Next
	}

	return vals
}

func removeNthLastNode(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	follower, leader := head, head
	for i := 0; i < n; i++ {
		leader = leader.Next
		// If n is the length of the linked list, then fast will be nil.
		// In this case, we want to remove the head of the linked list.
		if leader == nil {
			return head.Next
		}
	}

	for leader != nil && leader.Next != nil {
		follower = follower.Next
		leader = leader.Next
	}

	if follower.Next == nil {
		return nil
	}

	follower.Next = follower.Next.Next
	return head
}
