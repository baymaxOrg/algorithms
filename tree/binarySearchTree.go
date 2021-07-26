package tree

// BinaryTreeNode 定义二叉树的节点。
type BinaryTreeNode struct {
	Val   int
	Right *BinaryTreeNode
	Left  *BinaryTreeNode
}

// Insert 插入节点。
func Insert(root *BinaryTreeNode, val int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{
			Val:   val,
			Right: nil,
			Left:  nil,
		}
	}
	if root.Val > val {
		node := Insert(root.Left, val)
		root.Left = node
	} else {
		node := Insert(root.Right, val)
		root.Right = node
	}
	return root
}

// Create 将一个slice转成bst
func Create(list []int) *BinaryTreeNode {
	root := &BinaryTreeNode{
		Val:   list[0],
		Right: nil,
		Left:  nil,
	}
	for i := 1; i < len(list); i++ {
		Insert(root, list[i])
	}
	return root
}

// Search 查找
func Search(root *BinaryTreeNode, val int) bool {
	if root == nil {
		return false
	}
	if root.Val == val {
		return true
	}
	if root.Val > val {
		return Search(root.Left, val)
	} else {
		return Search(root.Right, val)
	}
}

func Delete(root *BinaryTreeNode, val int) *BinaryTreeNode {
	// 根节点为空
	if root == nil {
		return nil
	}
	// 查找需要删除的节点
	// 查找左树
	if root.Val > val {
		node := Delete(root.Left, val)
		root.Left = node
	} else if root.Val < val { // 朝朝作树。
		node := Delete(root.Right, val)
		root.Right = node
	} else { // 查找到了。
		// root := root
		// 左右节点都为空时。
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Right != nil { // 右子树存在时，找到右子树最小的。
			minRoot := root.Right
			parent := root
			for {
				//if minRoot == nil {
				//	break
				//}
				// 处理最小的节点。
				if minRoot.Left == nil { // 找到最小的节点
					parent.Left = minRoot.Right
					minRoot.Right = nil
					break
				}
				parent = minRoot
				minRoot = minRoot.Left
			}
			minRoot.Left = root.Left
			root.Left = nil
			minRoot.Right = root.Right
			root.Right = nil
			return minRoot
		} else { // 右子树不存在
			//if root.Right.Left、
			maxRoot := root.Left
			parent := root
			for {
				//if maxRoot == nil {
				//	break
				//}
				// 处理最小的节点。
				if maxRoot.Right == nil { // 找到最大的节点
					parent.Right = maxRoot.Left
					maxRoot.Left = nil
					break
				}
				parent = maxRoot
				maxRoot = maxRoot.Right
			}
			maxRoot.Left = root.Left
			root.Left = nil
			maxRoot.Right = root.Right
			root.Right = nil
			return maxRoot
		}
	}
	return root
	//return nil
}

func FindMaxElement(root *BinaryTreeNode) *BinaryTreeNode {
	var recursion func(r *BinaryTreeNode)
	var max *BinaryTreeNode

	recursion = func(r *BinaryTreeNode) {
		if r == nil {
			//max = nil
			return
		}
		if r.Right == nil {
			max = r
		}
		recursion(r.Right)
	}
	recursion(root)
	return max
}

func FindMinElement(root *BinaryTreeNode) *BinaryTreeNode {
	var recursion func(r *BinaryTreeNode)
	var min *BinaryTreeNode

	recursion = func(r *BinaryTreeNode) {
		if r == nil {
			//max = nil
			return
		}
		if r.Left == nil {
			min = r
		}
		recursion(r.Left)
	}
	recursion(root)
	return min
}
