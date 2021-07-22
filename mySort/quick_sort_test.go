package mysort_test

import (
	mysort "gitee.com/RocsSun/algorithms/mySort"
	"testing"
)

func TestQuickSort1(t *testing.T) {
	values := []int{4, 3, 2, 1, 0}
	mysort.QuickSort(values)
	for i, v := range values {
		if values[i] != v {
			t.Error("quick sort failed")
		}
	}
}
