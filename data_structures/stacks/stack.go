package stacks

type Queue[T any] struct {
	d1 Stack[T]
	d2 Stack[T]
}

func (q *Queue[T]) Enqueue(element T) {
	q.d1.Push(element)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.d2.IsEmpty() {
		for !q.d1.IsEmpty() {
			element, _ := q.d1.Pop()
			q.d2.Push(element)
		}
	}
	return q.d2.Pop()
}

type Stack[T any] struct {
	data []T
}

// Push method will push an element into the stack.
func (s *Stack[T]) Push(element T) {
	s.data = append(s.data, element)
}

// Pop method will pop an element from the stack.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element, true
}

// Peek method will return the top element of the stack.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// IsEmpty method will return true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func removeDuplicates(s string) string {
	stack := new(Stack[rune])
	for _, char := range s {
		if top, ok := stack.Peek(); ok && top == char {
			stack.Pop()
		} else {
			stack.Push(char)
		}
	}
	return string(stack.data)
}

func minRemoveParens(s string) string {
	type StackVal struct {
		index int
		val   rune
	}

	stack := new(Stack[StackVal])
	for i, char := range s {
		if char == '(' {
			stack.Push(StackVal{i, char})
		} else if char == ')' {
			if val, ok := stack.Peek(); ok && val.val == '(' {
				stack.Pop()
			} else {
				stack.Push(StackVal{i, char})
			}
		}
	}

	var skip Stack[int]
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		skip.Push(val.index)
	}

	if skip.IsEmpty() {
		return s
	}

	var result []rune
	for i, char := range s {
		if val, ok := skip.Peek(); ok && val == i {
			skip.Pop()
			continue
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}
