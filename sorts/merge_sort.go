package sorts

func MergeSort(nums []int) {
	merge_sort(nums, 0, len(nums)-1)
}

func merge_sort(nums []int, left, right int) {
	if left < right {
		mid := int((left + right) / 2)
		merge_sort(nums, left, mid)
		merge_sort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	i, j := left, mid+1
	var tmp []int
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, nums[i])
		i++
	}
	for j <= right {
		tmp = append(tmp, nums[j])
		j++
	}

	for t := left; t <= right; t++ {
		nums[t] = tmp[t-left]
	}
}
