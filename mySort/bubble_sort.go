package mysort

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	bubbleSort(nums)
}

// bubbleSort 冒泡排序的具体实现
func bubbleSort(nums []int) {
	// 控制便利的次数为slice的长度。
	for i := 0; i < len(nums); i++ {
		// 遍历无序区
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				// 交换元素。大的换后边，
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
}
