package sorts_test

import (
	"testing"

	mysort "gitee.com/RocsSun/algorithms/sorts"
)

func TestQuickSort1(t *testing.T) {
	values := []int{4, 3, 2, 1, 0}
	mysort.QuickSort(values)
	for i, v := range values {
		if values[i] != v {
			t.Error("quick sorts failed")
		}
	}
}
