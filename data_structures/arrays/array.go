package arrays

import "math"

func ProductOfOtherElements(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	res := make([]int, 0, len(arr))
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
	if len(arr) == 0 {
		return nil
	}

	res := make([]int, 0, len(arr))
	prefixProduct := make([]int, 0, len(arr))
	for _, v := range arr {
		if len(prefixProduct) == 0 {
			prefixProduct = append(prefixProduct, v)
		} else {
			prefixProduct = append(prefixProduct, prefixProduct[len(prefixProduct)-1]*v)
		}
	}

	suffixProduct := make([]int, 0, len(arr))
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

func SmallestWindowSorted(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}

	if len(arr) == 1 {
		return arr[0], arr[0]
	}

	right, left := 0, len(arr)-1
	minV, maxV := math.MaxInt, math.MinInt
	for i, val := range arr {
		maxV = max(maxV, val)
		if val < maxV {
			right = i
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		minV = min(minV, arr[i])
		if arr[i] > minV {
			left = i
		}
	}

	return left, right
}

func MaxSubarraySum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[1]
	}

	currSum, maxSum := 0, math.MinInt
	for _, v := range arr {
		currSum = max(v, v+currSum)
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}

func MaxSubarrayCircularSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[1]
	}

	// This could be optimized further by returning the sum from either MaxSubarraySum or MinSubarraySum.
	var sum int
	for _, v := range arr {
		sum += v
	}

	minV := MinSubarraySum(arr)
	circularSum := sum - minV
	return max(MaxSubarraySum(arr), circularSum)

}

func MinSubarraySum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[1]
	}

	currSum, minSum := 0, math.MaxInt
	for _, v := range arr {
		currSum = min(v, v+currSum)
		minSum = min(minSum, currSum)
	}

	return minSum
}
