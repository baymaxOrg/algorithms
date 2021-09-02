package sorts

// SelectSort Select sort.
// return sorted nums.
func SelectSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	return selectSort(nums)
}

// selectSort implement select sort.
func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}
