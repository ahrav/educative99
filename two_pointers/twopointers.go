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

func sortColors(colors []int) []int {
	if len(colors) <= 1 {
		return colors
	}

	red, white, blue := 0, 0, len(colors)-1
	for white <= blue {
		switch colors[white] {
		case 0:
			if colors[red] != 0 {
				colors[red], colors[white] = colors[white], colors[red]
			}
			red++
			white++
		case 1:
			white++
		default:
			if colors[blue] != 2 {
				colors[white], colors[blue] = colors[blue], colors[white]
			}
			blue--
		}
	}

	return colors
}

// This approach is slower than the two pointer approach.
// It also uses more memory.
func isHappy(n int) bool {
	if n == 0 {
		return false
	}

	seen := make(map[int]struct{})
	for n != 1 {
		if _, ok := seen[n]; ok {
			return false
		}

		seen[n] = struct{}{}
		n = sumOfSquares(n)
	}

	return true
}

func isHappyTwoPointer(n int) bool {
	if n == 0 {
		return false
	}

	slow, fast := n, sumOfSquares(n)
	for slow != fast {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))
	}

	return slow == 1
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}

	return sum
}

func containsCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}

	return false
}
