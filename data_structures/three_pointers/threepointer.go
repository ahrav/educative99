package three_pointers

// Three pointer...
func mergeSortedSlices(nums1 []int, m int, nums2 []int, n int) {
	endNums1Idx, endNums2Idx, lastMergeIdx := m-1, n-1, m+n-1

	for endNums1Idx >= 0 && endNums2Idx >= 0 {
		if nums2[endNums2Idx] > nums1[endNums1Idx] {
			nums1[lastMergeIdx] = nums2[endNums2Idx]
			endNums2Idx--
		} else {
			nums1[lastMergeIdx] = nums1[endNums1Idx]
			endNums1Idx--
		}
		lastMergeIdx--
	}

	for endNums2Idx >= 0 {
		nums1[lastMergeIdx] = nums2[endNums2Idx]
		endNums2Idx--
		lastMergeIdx--
	}
}
