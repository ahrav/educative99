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
