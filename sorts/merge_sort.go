package sorts

// MergeSort function
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

// mergeSort recursive make subslice order, and merge.
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	middle := len(nums) / 2
	nums1 := mergeSort(nums[:middle])
	nums2 := mergeSort(nums[middle:])
	return merge(nums1, nums2)
}

// merge merge two order silce
func merge(nums1, nums2 []int) []int {
	var i = 0
	var j = 0
	var tmp []int = make([]int, len(nums1)+len(nums2))

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			tmp[i+j] = nums1[i]
			i++
		} else {
			tmp[i+j] = nums2[j]
			j++
		}
	}

	for i < len(nums1) {
		tmp[i+j] = nums1[i]
		i++
	}

	for j < len(nums2) {
		tmp[i+j] = nums1[j]
		j++
	}
	return tmp
}
