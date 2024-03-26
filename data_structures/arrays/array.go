package arrays

func ProductOfOtherElements(arr []int) []int {
	var res []int
	if len(arr) == 0 {
		return res
	}

	total := 1
	for _, v := range arr {
		total *= v
	}

	for _, v := range arr {
		res = append(res, total/v)
	}

	return res
}

func ProductOfOtherElementsNoDivide(arr []int) []int {
	var res []int
	if len(arr) == 0 {
		return res
	}

	var prefixProduct []int
	for _, v := range arr {
		if len(prefixProduct) == 0 {
			prefixProduct = append(prefixProduct, v)
		} else {
			prefixProduct = append(prefixProduct, prefixProduct[len(prefixProduct)-1]*v)
		}
	}

	var suffixProduct []int
	for i := len(arr) - 1; i >= 0; i-- {
		if len(suffixProduct) == 0 {
			suffixProduct = append(suffixProduct, arr[i])
		} else {
			suffixProduct = append([]int{suffixProduct[0] * arr[i]}, suffixProduct...)
		}
	}

	for i := range arr {
		if i == 0 {
			res = append(res, suffixProduct[0])
		} else if i == len(arr)-1 {
			res = append(res, prefixProduct[i-1])
		} else {
			res = append(res, suffixProduct[i+1]*prefixProduct[i-1])
		}
	}

	return res
}
