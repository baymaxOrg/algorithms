package sorts

// QuickSort 快排函数
func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	return quickSort(nums)
}

// quickSort 预处理
func quickSort(nums []int) []int {
	_quickSort(nums, 0, len(nums)-1)
	return nums
}

// _quickSort 快排
func _quickSort(nums []int, left, right int) {
	if left < right {
		middle := partition(nums, left, right)
		_quickSort(nums, left, middle)
		_quickSort(nums, middle+1, right)
	}
}

// partition 移动元素。
func partition(nums []int, left, right int) int {
	tmp := nums[left]
	for right > left {
		for right > left && tmp <= nums[right] {
			right--
		}
		nums[left] = nums[right]
		for right > left && tmp >= nums[left] {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = tmp
	return left
}
