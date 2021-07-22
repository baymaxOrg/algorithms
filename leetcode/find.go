package leetcode

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	target := len(nums1) + len(nums2)
	if target%2 == 1 {
		k := int(target / 2)
		return float64(findK(nums1, nums2, k))
	} else {
		k1, k2 := int(target/2-1), int(target/2)
		fmt.Println(k1, k2)
		a := findK(nums1, nums2, k1)
		b := findK(nums1, nums2, k2)
		fmt.Println(a, b)
		return float64((a + b)) / 2.0
	}
}

func findK(nums1, nums2 []int, k int) int {
	i1, i2 := 0, 0
	tmp := -1
	for i := 0; i <= k; i++ {
		if i1 < len(nums1) && i2 < len(nums2) {
			if nums1[i1] <= nums2[i2] {
				tmp = nums1[i1]
				i1++
			} else {
				tmp = nums2[i2]
				i2++
			}
		} else if i1 < len(nums1) {
			tmp = nums1[i1]
			i1++
		} else {
			tmp = nums2[i2]
			i2++
		}

	}
	return tmp
}
