package arrays

import (
	"math"
	"sort"
)

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

type SortedList struct {
	data []int
}

func (s *SortedList) Add(value int) int {
	index := sort.Search(len(s.data), func(i int) bool {
		return s.data[i] >= value
	})
	s.data = append(s.data, 0)
	copy(s.data[index+1:], s.data[index:])
	s.data[index] = value
	return index
}

func (s *SortedList) Get() []int {
	return s.data
}

func SmallerCounts(arr []int) []int {
	result := make([]int, 0, len(arr))

	var sortedLst SortedList
	for i := len(arr) - 1; i >= 0; i-- {
		result = append([]int{sortedLst.Add(arr[i])}, result...)
	}

	return result
}

func minRemovals(arr []int, k int, threshold int) int {
	sort.Ints(arr) // Sort the array in ascending order
	n := len(arr)
	windowSum := sum(arr[:k]) // Sum of the first k elements
	removals := 0

	if windowSum > threshold {
		return 1 // No valid subset of k elements exists
	}

	var include bool
	for i := k; i < n; i++ {
		if !include {
			windowSum += arr[i] - arr[i-k] // Slide the window
		} else {
			windowSum += arr[i]
			include = false
		}
		if windowSum > threshold {
			removals++
			windowSum -= arr[i] // Remove the current element
			include = true
		}
	}

	return removals
}

func sum(arr []int) int {
	result := 0
	for _, num := range arr {
		result += num
	}
	return result
}
