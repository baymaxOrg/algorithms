package sorts_test

import (
	"testing"

	ms "gitee.com/RocsSun/algorithms/sorts"
)

func TestMergeSort(t *testing.T) {
	values := []int{4, 3, 2, 1, 0}
	ms.MergeSort(values)
	for i, v := range values {
		if values[i] != v {
			t.Error("merge sorts failed")
		}
	}
}

func TestMergeSort1(t *testing.T) {
	values := []int{4, 2, 2, 1, 0}
	ms.MergeSort(values)
	if values[0] != 0 && values[1] != 1 && values[2] != 2 && values[3] != 2 && values[0] != 4 {
		t.Error("merge sorts failed")
	}
}
