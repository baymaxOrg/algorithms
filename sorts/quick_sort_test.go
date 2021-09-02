package sorts_test

import (
	"testing"

	ms "gitee.com/RocsSun/algorithms/sorts"
)

func TestQuickSort(t *testing.T) {
	values := []int{4, 3, 2, 2, 0}
	ms.QuickSort(values)
	if values[0] != 0 && values[1] != 2 && values[2] != 2 && values[3] != 3 && values[4] != 4 {
		t.Error("bubble sorts failed")
	}
}

func TestQuickSort1(t *testing.T) {
	values := []int{4, 3, 2, 1, 0}
	ms.QuickSort(values)
	for i, v := range values {
		if values[i] != v {
			t.Error("quick sorts failed")
		}
	}
}
