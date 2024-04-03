package strings

import "sort"

func AnagramIndices(w, s string) []int {
	if len(w) > len(s) {
		return nil
	}

	var target int32
	for _, c := range w {
		target += c
	}

	var result []int
	left, right := 0, len(w)
	var curr int32
	for _, v := range s[:right] {
		curr += v
	}

	for right <= len(s) {
		if curr == target {
			result = append(result, left)
		}

		if right < len(s) {
			curr += int32(s[right])
		}
		curr -= int32(s[left])

		left++
		right++
	}

	return result
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
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

func ReverseString(s string) string {
	// Convert string to a slice of runes.
	runes := []rune(s)

	// Reverse the runes in place.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string and return.
	return string(runes)
}

func IsPalindromePairs(arr []string) [][]int {
	m := make(map[string]int)
	for i, v := range arr {
		m[v] = i
	}

	var result [][]int
	for i, v := range arr {
		for j := range v {
			pre, suf := v[j:], v[:j]
			revPre := ReverseString(pre)
			revSuf := ReverseString(suf)

			if val, ok := m[revSuf]; ok && isPalindrome(pre) {
				if val != i {
					result = append(result, []int{i, val})
				}
			}

			if val, ok := m[revPre]; ok && isPalindrome(suf) {
				if val != i {
					result = append(result, []int{i, val})
				}
			}
		}
	}

	return result
}

func HashUniqueASCIIChars(s string) bool {
	if len(s) > 128 {
		return false
	}

	var vector [2]uint64
	for _, c := range s {
		idx := c / 64
		bitPos := c % 64

		bitMask := uint64(1 << bitPos)
		if vector[idx]&bitMask != 0 {
			return false
		}

		vector[idx] |= bitMask
	}

	return true
}

func HashUniqueUnicodeSort(s string) bool {
	sortString := func(s string) string {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		return string(runes)
	}

	sortString(s)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}

	return true
}

func HasUniqueUnicodeHash(s string) bool {
	m := make(map[int32]struct{}, len(s))
	for _, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = struct{}{}
		} else {
			return false
		}
	}

	return true
}
