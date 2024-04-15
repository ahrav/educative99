package strings

import (
	"sort"
	"strconv"
	"strings"
)

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

func CheckPermutationUnicode(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[int32]int)
	for _, c := range a {
		m[c]++
	}

	for _, c := range b {
		if _, ok := m[c]; !ok {
			return false
		} else {
			m[c]--
		}
	}

	return true
}

func CheckPermutationASCII(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	var count [128]int
	for _, c := range a {
		count[c]++
	}

	for _, c := range b {
		if count[c] == 0 {
			return false
		}
		count[c]--
	}

	return true
}

func URLify(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if c == ' ' {
			sb.WriteString("%20")
			continue
		}
		sb.WriteRune(c)
	}

	return sb.String()
}

func PalindromePermutation(s string) bool {
	if len(s) < 2 {
		return true
	}

	m := make(map[int32]int)
	for _, c := range s {
		if c == ' ' {
			continue
		}
		m[c]++
	}

	if len(m) == 1 {
		return true
	}

	var oddCnt int
	for _, v := range m {
		if v%2 != 0 {
			oddCnt++
			if oddCnt > 1 {
				return false
			}
		}
	}

	return true
}

func PalindromePermutationVector(s string) bool {
	var bitVector uint8
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			idx := c - 'a'
			bitVector ^= 1 << idx
		}
	}

	return (bitVector & (bitVector - 1)) == 0
}

func OneAway(src, target string) bool {
	srcLen, tgtLen := len(src), len(target)
	if srcLen-tgtLen > 1 {
		return false
	}

	var longer, shorter string
	if srcLen > tgtLen {
		longer = src
		shorter = target
	} else {
		longer = target
		shorter = src
	}

	var idx int
	var hasDiff bool
	for i := range longer {
		if longer[i] == shorter[idx] {
			idx++
			continue
		}
		if hasDiff {
			return false
		}
		hasDiff = true

		if srcLen == tgtLen {
			idx++
		}
	}

	return true
}

func StringCompressions(s string) string {
	compLen := countCompression(s)
	if compLen == len(s) {
		return s
	}

	var sb strings.Builder
	sb.Grow(compLen)
	cnt := 1
	char := s[0]

	sb.WriteByte(char)
	for i := 1; i < len(s); i++ {
		if s[i] != char {
			char = s[i]
			sb.WriteString(strconv.Itoa(cnt))
			sb.WriteByte(char)
			cnt = 1
			continue
		}
		cnt++
	}
	sb.WriteString(strconv.Itoa(cnt))

	return sb.String()
}

func countCompression(s string) int {
	var compLen, count int
	for i := 0; i < len(s); i++ {
		count++
		if i+1 >= len(s) || s[i] != s[i+1] {
			compLen += 1 + len(strconv.Itoa(compLen))
			count = 0
		}
	}

	return compLen
}

// TODO: add test
func RotateMatrix(matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return false
	}

	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]

			matrix[first][i] = matrix[last-offset][first]

			matrix[last-offset][first] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}

	return true
}

func ZeroMatrix(matrix [][]int) {
	var rowIdx, colIdx int
exit:
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rowIdx, colIdx = i, j
				break exit
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		matrix[rowIdx][i] = 0
		matrix[i][colIdx] = 0
	}
}

func IsRotationSubstring(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	combined := a + a
	return strings.Contains(combined, b)
}
