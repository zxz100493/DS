package binarysearch

import "fmt"

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

type BinarySearchTreeNode struct {
	Value int64
	Times int64
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

func (tree *BinarySearchTree) Add(value int64) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}

	tree.Root.Add(value)
}

/*
如果添加元素时是棵空树，那么初始化根节点。

然后添加的值和根节点比较，判断是要插入到根节点左子树还是右子树，还是不用插入。

当值比根节点小时，元素要插入到根节点的左子树中，当值比根节点大时，元素要插入到根节点的右子树中，相等时不插入，只更新次数。

然后再分别对根节点的左子树和右子树进行递归操作即可。
*/
func (node *BinarySearchTreeNode) Add(value int64) {
	if value < node.Value {
		if node.Left != nil {
			node.Left = &BinarySearchTreeNode{Value: value}
		} else {
			node.Left.Add(value)
		}
	} else if value > node.Value {
		if node.Right != nil {
			node.Right = &BinarySearchTreeNode{Value: value}
		} else {
			node.Right.Add(value)
		}
	} else {
		node.Times = node.Times + 1
	}
}

// find max value
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	if node.Left == nil {
		return node
	}
	return node.Left.FindMinValue()
}
func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	if node.Right == nil {
		return node
	}
	return node.Right.FindMinValue()
}

func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Find(value)
}

/*
	如果是空树，返回 nil，否则与根节点比较。

如果刚刚好等于根节点的值，返回该节点，否则根据值的比较，继续往左子树或右字树递归查找。
*/
func (node *BinarySearchTreeNode) Find(value int64) *BinarySearchTreeNode {
	if node.Value == value {
		return node
	} else if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			return nil
		}
		return node.Left.Find(value)
	} else {
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(value)
	}
}

// findParent
func (tree *BinarySearchTree) FindParent(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		// / 如果是空树，返回空
		return nil
	}
	// 如果根节点等于该值，根节点其没有父节点，返回nil
	if tree.Root.Value == value {
		return nil
	}
	return tree.Root.FindParent(value)
}

func (node *BinarySearchTreeNode) FindParent(value int64) *BinarySearchTreeNode {
	if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		leftTree := node.Left
		if leftTree == nil {
			return nil
		}

		if leftTree.Value == value {
			return node
		} else {
			return leftTree.FindParent(value)
		}
	} else {
		rightTree := node.Right
		if rightTree == nil {
			return nil
		}

		if rightTree.Value == value {
			return node
		} else {
			return rightTree.FindParent(value)
		}
	}

}

/*
	删除元素有四种情况：

第一种情况，删除的是根节点，且根节点没有儿子，直接删除即可。
第二种情况，删除的节点有父亲节点，但没有子树，也就是删除的是叶子节点，直接删除即可。
第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点，这时二叉查找树的性质又满足了。右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到。
第四种情况，删除的节点只有一个子树，那么该子树直接替换被删除的节点即可。
*/
func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		return
	}
	// 查找该值是否存在
	node := tree.Root.Find(value)

	if node == nil {
		return
	}

	parent := tree.Root.FindParent(value)
	// 第一种情况，删除的是根节点，且根节点没有儿子
	if parent == nil && node.Left == nil && node.Right == nil {
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		// 第二种情况，删除的节点有父亲节点，但没有子树
		// 如果删除的是节点是父亲的左儿子，直接将该值删除即可
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			parent.Right = nil // 删除的原来是父亲的右儿子，直接将该值删除即可
		}
	} else if node.Left != nil && node.Right != nil {
		// 第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点。
		// 右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到，替换后二叉查找树的性质又满足了。

		// 找右子树中最小的值，一直往右子树的左边找
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		tree.Delete(minNode.Value) // 把最小的节点删掉

		// 最小值的节点替换被删除节点
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		// 第四种情况，只有一个子树，那么该子树直接替换被删除的节点即可

		// 父亲为空，表示删除的是根节点，替换树根
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}

		// 左子树不为空
		if node.Left != nil {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的左子树接班
			if parent.Left != nil && value == node.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的右子树接班
			if parent.Left != nil && value == node.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}

// 中序遍历

func (tree *BinarySearchTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *BinarySearchTreeNode) MidOrder() {
	if node == nil {
		return
	}

	node.Left.MidOrder()
	for i := 0; i < int(node.Times); i++ {
		fmt.Println(node.Value)
	}

	node.Right.MidOrder()
}
