package utils_test

import (
	"gitee.com/RocsSun/algorithms/utils"
	"testing"
)

func TestShuffleSlice(t *testing.T) {
	var s = []int{1, 2, 3, 4, 5, 6}
	utils.ShuffleSlice(s)
	if !(s[0] != 1 || s[1] != 2 || s[2] != 3 || s[3] != 4 || s[4] != 5 || s[5] != 6) {
		t.Error("shuffle slice failed!")
	}
}

func TestShuffleSlice2(t *testing.T) {
	var s []int
	for i := 0; i < 10000; i ++ {
		s = append(s, i)
	}
	utils.ShuffleSlice(s)
	win := true
	for i := 0; i < len(s); i++ {
		if s[i] != i {
			win = false
		}
	}
	if win {
		t.Error("shuffle slice failed!")
	}
}