package mysort

func HeapSort(nums []int) {
	BuildMinHeap(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		siftDown(nums, 0, i-1)
	}
}

func BuildMinHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, i, len(nums)-1)
	}
}

func siftDown(nums []int, root, lastLeaf int) {
	tmp := nums[root]
	i := root
	j := 2*i + 1
	for j <= lastLeaf {
		// 将j指向较小的孩子节点。
		if j+1 <= lastLeaf && nums[j] > nums[j+1] {
			j++
		}
		// 比较子树的根节点与较小孩子的值，
		if tmp > nums[j] {
			nums[i] = nums[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	nums[i] = tmp
}
