package strings

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

func AnagramIndicesMap(w, s string) []int {
	if len(w) > len(s) {
		return nil
	}

	freq := make(map[int32]int)
	for _, c := range w {
		freq[c]++
	}

	for _, c := range s {
		freq[c]--
		if freq[c] == 0 {
			delete(freq, c)
		}
	}

	var result []int

	if len(freq) == 0 {
		result = append(result, 0)
	}

	for i := len(w); i < len(s); i++ {
		left, right := s[i-len(w)], s[i]
		freq[int32(left)]++
		if freq[int32(left)] == 0 {
			delete(freq, int32(left))
		}

		freq[int32(right)]--
		if freq[int32(right)] == 0 {
			delete(freq, int32(right))
		}

		if len(freq) == 0 {
			result = append(result, i-len(w)+1)
		}
	}

	return result
}
