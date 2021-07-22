package mysort_test

import (
	mysort "gitee.com/RocsSun/algorithms/mySort"
	"gitee.com/RocsSun/algorithms/utils"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	values := []int{4, 3, 2, 1, 0}
	utils.ShuffleSlice(values)
	mysort.BubbleSort(values)
	if values[0] != 0 && values[1] != 1 && values[2] != 2 && values[3] != 3 && values[4] != 4 {
		t.Error("bubble sort failed")
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{4, 3, 2, 2, 0}
	mysort.BubbleSort(values)
	if values[0] != 0 && values[1] != 2 && values[2] != 2 && values[3] != 3 && values[4] != 4 {
		t.Error("bubble sort failed")
	}
}