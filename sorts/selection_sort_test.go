package sorts_test

import (
	ms "gitee.com/RocsSun/algorithms/sorts"
	"testing"
)

func TestSelectSort(t *testing.T) {
	values := []int{4, 3, 2, 2, 0}
	ms.SelectSort(values)
	if values[0] != 0 && values[1] != 2 && values[2] != 2 && values[3] != 3 && values[4] != 4 {
		t.Error("bubble sorts failed")
	}
}

func TestSelectSort1(t *testing.T) {
	values := []int{4, 3, 1, 2, 0}
	ms.SelectSort(values)
	if values[0] != 0 && values[1] != 1 && values[2] != 2 && values[3] != 3 && values[4] != 4 {
		t.Error("bubble sorts failed")
	}
}