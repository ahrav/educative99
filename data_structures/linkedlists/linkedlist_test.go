package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	// Test case 1: Empty linked list
	emptyList := &EduLinkedList{}
	reversedEmptyList := reverse(emptyList.head)
	assert.Nil(t, reversedEmptyList)

	// Test case 2: Linked list with a single node
	singleNodeList := &EduLinkedList{}
	singleNodeList.CreateLinkedList([]int{1})
	reversedSingleNodeList := reverse(singleNodeList.head)
	if reversedSingleNodeList.data != 1 || reversedSingleNodeList.next != nil {
		t.Errorf("Expected {1, nil}, but got {%d, %v}", reversedSingleNodeList.data, reversedSingleNodeList.next)
	}

	// Test case 3: Linked list with multiple nodes
	list := &EduLinkedList{}
	list.CreateLinkedList([]int{1, 2, 3, 4, 5})
	reversedList := reverse(list.head)
	expectedData := []int{5, 4, 3, 2, 1}
	temp := reversedList
	i := 0
	for temp != nil {
		if temp.data != expectedData[i] {
			t.Errorf("Expected %v, but got %v", expectedData, getLinkedListData(reversedList))
			break
		}
		temp = temp.next
		i++
	}
}

func getLinkedListData(head *EduLinkedListNode) []int {
	data := []int{}
	temp := head
	for temp != nil {
		data = append(data, temp.data)
		temp = temp.next
	}
	return data
}

func TestAlternateValues(t *testing.T) {
	list := &EduLinkedList{}
	list.CreateLinkedList([]int{1, 2, 3, 4, 5})
	AlternateValues(list.head)
	expectedData := []int{1, 3, 2, 5, 4}
	i := 0
	tmp := list.head
	for tmp != nil {
		assert.Equal(t, expectedData[i], tmp.data)
		tmp = tmp.next
		i++
	}
}
