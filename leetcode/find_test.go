package leetcode

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	a := FindMedianSortedArrays(nums1, nums2)
	fmt.Println(a)
}
