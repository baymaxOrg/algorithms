package sorts_test

import (
	"testing"

	ms "gitee.com/RocsSun/algorithms/sorts"
	"gitee.com/RocsSun/algorithms/utils"
)

func TestBubbleSort(t *testing.T) {
	values := []int{4, 3, 2, 1, 0}
	utils.ShuffleSlice(values)
	ms.BubbleSort(values)
	if values[0] != 0 && values[1] != 1 && values[2] != 2 && values[3] != 3 && values[4] != 4 {
		t.Error("bubble sorts failed")
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{4, 3, 2, 2, 0}
	ms.BubbleSort(values)
	if values[0] != 0 && values[1] != 2 && values[2] != 2 && values[3] != 3 && values[4] != 4 {
		t.Error("bubble sorts failed")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	var nums []int
	for i := 0; i < 10000000; i ++ {
		nums = append(nums, i)
	}
	b.ResetTimer()
	ms.BubbleSort(nums)
}