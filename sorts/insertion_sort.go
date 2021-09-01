//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func InsertionSort(nums []int) []int {
	return insertionSort(nums)
}

func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i ++ {
		tmp := nums[i]
		j := i - 1
		for j >= 0 && tmp < nums[j] {
			//if
			nums[j], nums[j+1] = nums[j+1], nums[j]
			j--
		}
	}
	return nums
}