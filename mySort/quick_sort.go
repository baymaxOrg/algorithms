package mysort

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left < right {
		position := partition(nums, left, right)
		quickSort(nums, left, position)
		quickSort(nums, position+1, right)
	}
}

func partition(nums []int, left, right int) int {
	tmp := nums[left]
	for left < right {
		for left < right && tmp <= nums[right] {
			right--
		}
		nums[left] = nums[right]

		for left < right && tmp >= nums[left] {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = tmp
	return left
}

func _quickSort(nums []int, left, right int) {
	i, j := left, right
	if i < j {
		tmp := nums[i]
		for i < j {
			for i < right && tmp <= nums[j] {
				j--
			}
			nums[i] = nums[j]

			for i < j && tmp >= nums[i] {
				i++
			}
			nums[j] = nums[i]
		}
		nums[i] = tmp
		_quickSort(nums, left, i)
		_quickSort(nums, i+1, right)
	}
}
