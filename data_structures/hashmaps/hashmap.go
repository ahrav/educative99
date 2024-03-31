package hashmaps

type RequestLogger struct {
	m     map[string]int
	limit int
}

func (l *RequestLogger) requestLoggerInit(timeLimit int) {
	l.m = make(map[string]int)
	l.limit = timeLimit
}

func (l *RequestLogger) messageRequestDecision(timestamp int, request string) bool {
	val, ok := l.m[request]
	if !ok {
		l.m[request] = timestamp
		return true
	}

	if timestamp-val >= l.limit {
		l.m[request] = timestamp
		return true
	}

	return false
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, len(nums2))
	m := make(map[int]int)

	for _, current := range nums2 {
		for len(stack) > 0 && current > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = current
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, current)
	}

	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		if val, ok := m[num]; ok {
			ans[i] = val
			continue
		}
		ans[i] = -1
	}

	return ans
}

func isIsomorphic(string1 string, string2 string) bool {
	if len(string1) != len(string2) {
		return false
	}

	m := make(map[rune]rune)
	for i, c := range string1 {
		v, ok := m[c]
		if !ok {
			x := rune(string2[i])
			if val, ok := m[x]; ok {
				if val == x {
					return false
				}
			}
			m[c] = x
			continue
		}

		if v != rune(string2[i]) {
			return false
		}
	}

	return true
}
