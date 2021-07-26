package tree_test

import (
	"fmt"
	"gitee.com/RocsSun/algorithms/tree"
	"gitee.com/RocsSun/algorithms/utils"
	"testing"
)

func TestCreate(t *testing.T) {
	var a = []int{2, 1, 4, 0, 3}
	root := tree.Create(a)
	if !(root.Val == 2 && root.Left.Val == 1 && root.Left.Left.Val == 0 && root.Right.Val == 4 && root.Right.Left.Val == 3) {
		t.Error("insert binary search tree failed!")
	}
}

func TestSearch(t *testing.T) {
	var a = []int{2, 1, 4, 0, 3}
	root := tree.Create(a)
	if tree.Search(root, 5) {
		t.Error("bst shouldn't include 5")
	}
	if !tree.Search(root, 1) {
		t.Error("bst should include 1")
	}
}

func TestInsert(t *testing.T) {
	var a = []int{2, 1, 4, 0, 3}
	root := tree.Create(a)
	if !(root.Val == 2 && root.Left.Val == 1 && root.Left.Left.Val == 0 && root.Right.Val == 4 && root.Right.Left.Val == 3) {
		t.Error("insert binary search tree failed!")
	}
}

func TestFindMaxElement(t *testing.T) {
	var a = []int{2, 1, 4, 0, 3}
	root := tree.Create(a)
	if !(root.Val == 2 && root.Left.Val == 1 && root.Left.Left.Val == 0 && root.Right.Val == 4 && root.Right.Left.Val == 3) {
		t.Error("insert binary search tree failed!")
	}
	max := tree.FindMaxElement(root)
	if max.Val != 4 {
		t.Error("find max element failed!")
	}
}

func TestFindMinElement(t *testing.T) {
	var a = []int{2, 1, 4, 0, 3}
	root := tree.Create(a)
	max := tree.FindMinElement(root)
	if max.Val != 0 {
		t.Error("find max element failed!")
	}
}

func TestFindMaxElement2(t *testing.T) {
	var a []int
	for i := 0; i < 10000; i++ {
		a = append(a, i)
	}
	utils.ShuffleSlice(a)
	root := tree.Create(a)
	max := tree.FindMaxElement(root)
	if max.Val != 9999 {
		t.Error("find max element failed!")
	}
}

func TestDelete(t *testing.T) {
	var a = []int{2, 1, 4, 0, 3}
	root := tree.Create(a)
	tree.Delete(root, 4)
	fmt.Println(root)
	var s int8
	s = -1
	fmt.Println(s)
	s = s << 8
	fmt.Println(s)
}