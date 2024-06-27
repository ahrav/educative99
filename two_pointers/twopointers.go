package two_pointers

import (
	"strconv"
	"strings"
)

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

func getMiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func circularArrayLoop(nums []int) bool {
	size := len(nums)
	if size == 1 {
		return true
	}

	for i := range nums {
		slow, fast := i, i

		forward := nums[i] > 0

		for {
			slow = nextStep(slow, nums[slow], size)
			if isNotCycle(nums, forward, slow) {
				break
			}

			fast = nextStep(fast, nums[fast], size)
			if isNotCycle(nums, forward, fast) {
				break
			}

			fast = nextStep(fast, nums[fast], size)
			if isNotCycle(nums, forward, fast) {
				break
			}

			if slow == fast {
				return true
			}
		}
	}

	return false
}

func nextStep(pointer, val, size int) int {
	result := (pointer + val) % size
	if result < 0 {
		result += size
	}
	return result
}

func isNotCycle(nums []int, prevDirection bool, pointer int) bool {
	currDirection := nums[pointer] >= 0

	if (prevDirection != currDirection) || (abs(nums[pointer])%len(nums) == 0) {
		return true
	} else {
		return false
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}

func isBadVersion(n int) bool {
	return n >= 4
}

// No test for this one :(
// Weird edudative test.
func firstBadVersion(n int) (int, int) {
	if n == 1 {
		return 1, 1
	}

	count := 0
	left, right := 1, n
	for left < right {
		middle := (left + right) / 2
		count++
		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle + 1
		}

	}

	return left, count
}

func binarySearchRotated(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}

		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func addBinary(str1, str2 string) string {
	if str1 == "" {
		return str2
	}

	if str2 == "" {
		return str1
	}

	carry := 0
	var result strings.Builder
	for i, j := len(str1)-1, len(str2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += int(str1[i] - '0')
		}

		if j >= 0 {
			sum += int(str2[j] - '0')
		}

		result.WriteString(strconv.Itoa(sum % 2))
		carry = sum / 2
	}

	return reverse(result.String())
}

func reverse(str string) string {
	var result strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}

	return result.String()
}
