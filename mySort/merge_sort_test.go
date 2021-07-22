package mysort_test

import (
	mysort "gitee.com/RocsSun/algorithms/mySort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	values := []int{4, 3, 2, 1, 0}
	mysort.MergeSort(values)
	for i, v := range values {
		if values[i] != v {
			t.Error("merge sort failed")
		}
	}
}

func TestMergeSort1(t *testing.T) {
	values := []int{4, 2, 2, 1, 0}
	mysort.MergeSort(values)
	if values[0] != 0 && values[1] != 1 && values[2] != 2 && values[3] != 2 && values[0] != 4 {
		t.Error("merge sort failed")
	}
}
