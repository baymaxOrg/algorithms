package tree

// BinaryTreeNode 定义二叉树的节点。
type BinaryTreeNode struct {
	Val   int
	Right *BinaryTreeNode
	Left  *BinaryTreeNode
}

//func createTree(nums []int) *BinaryTreeNode {
//	rand.Rand()
//}
