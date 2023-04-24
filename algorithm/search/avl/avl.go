package avl

type AVLTree struct {
	Root *AVLTreeNode
}

// 其中 Height 表示以该节点作为树的根节点时该树的高度，方便计算平衡因子。
type AVLTreeNode struct {
	Value  int64
	Times  int64
	Height int64
	Left   *AVLTreeNode
	Right  *AVLTreeNode
}

func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

// get the height
func (node *AVLTreeNode) UpdateHeight() {
	if node == nil {
		return
	}

	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	maxHeight := leftHeight
	if rightHeight > leftHeight {
		maxHeight = rightHeight
	}
	node.Height = maxHeight + 1
}

// 计算树的平衡因子，也就是左右子树的高度差，代码如下：
func (node *AVLTreeNode) BalanceFactor() int64 {
	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return leftHeight - rightHeight
}

// 单右旋操作
func RightRotation(Root *AVLTreeNode) *AVLTreeNode {
	Pivot := Root.Left
	B := Pivot.Right
	Pivot.Right = Root
	Root.Left = B

	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

// 单左旋操作
func LeftRotation(Root *AVLTreeNode) *AVLTreeNode {
	Pivot := Root.Right
	B := Pivot.Left
	Pivot.Left = Root
	Root.Right = B

	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

// 先左后右旋操作
func LeftRightRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Left = LeftRotation(node.Left)
	return RightRotation(node)
}

// 先右后左旋操作，
func RightLeftRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Right = RightRotation(node.Left)
	return LeftRotation(node)
}

func (tree *AVLTree) Add(value int64) {
	tree.Root = tree.Root.Add(value)
}

func (node *AVLTreeNode) Add(value int64) *AVLTreeNode {
	// 添加值到根节点node，如果node为空，那么让值成为新的根节点，树的高度为1
	if node == nil {
		return &AVLTreeNode{Value: value, Height: 1}
	}
	// 如果值重复，什么都不用做，直接更新次数
	if node.Value == value {
		node.Times = node.Times + 1
		return node
	}
	// 辅助变量
	var newTreeNode *AVLTreeNode

	if value > node.Value {
		// 插入的值大于节点值，要从右子树继续插入
		node.Right = node.Right.Add(value)
		// 平衡因子，插入右子树后，要确保树根左子树的高度不能比右子树低一层。
		factor := node.BalanceFactor()
		// 右子树的高度变高了，导致左子树-右子树的高度从-1变成了-2。
		if factor == -2 {
			if value > node.Right.Value {
				newTreeNode = LeftRotation(node)
			} else {
				newTreeNode = RightRotation(node)
			}
		}
	} else {
		// 插入的值小于节点值，要从左子树继续插入
		node.Left = node.Left.Add(value)
		// 平衡因子，插入左子树后，要确保树根左子树的高度不能比右子树高一层。
		factor := node.BalanceFactor()
		// 左子树的高度变高了，导致左子树-右子树的高度从1变成了2。
		if factor == 2 {
			if value < node.Left.Value {
				newTreeNode = RightRotation(node) // 表示在左子树上插上左儿子导致失衡，需要单右旋：
			} else {
				newTreeNode = LeftRightRotation(node) //表示在左子树上插上右儿子导致失衡，先左后右旋：
			}
		}
	}

	if newTreeNode == nil {
		node.UpdateHeight() // 表示什么旋转都没有，根节点没变，直接刷新树高度
		return node
	} else {
		newTreeNode.UpdateHeight() // 旋转了，树根节点变了，需要刷新新的树根高度
		return newTreeNode
	}
}

func (node *AVLTreeNode) Delete(value int64) *AVLTreeNode {
	if node == nil {
		return nil // 如果是空树，直接返回
	}

	if value < node.Value {
		node.Left = node.Left.Delete(value) // 从左子树开始删除
		node.Left.UpdateHeight()            // 删除后要更新该子树高度
	} else if value > node.Value {
		node.Right = node.Right.Delete(value) // 从右子树开始删除
		node.Right.UpdateHeight()             // 删除后要更新该子树高度
	} else {
		// 找到该值对应的节点
		// 该节点没有左右子树
		// 第一种情况，删除的节点没有儿子，直接删除即可。
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// 该节点有两棵子树，选择更高的哪个来替换
		// 第二种情况，删除的节点下有两个子树，选择高度更高的子树下的节点来替换被删除的节点，如果左子树更高，选择左子树中最大的节点，
		// 也就是左子树最右边的叶子节点，如果右子树更高，选择右子树中最小的节点，也就是右子树最左边的叶子节点。最后，删除这个叶子节点。
		if node.Left != nil && node.Right != nil {
			// 左子树更高，拿左子树中最大值的节点替换
			if node.Left.Height > node.Right.Height {
				maxNode := node.Left
				for maxNode.Right != nil {
					maxNode = maxNode.Right
				}
				// 最大值的节点替换被删除节点
				node.Value = maxNode.Value
				node.Times = maxNode.Times

				node.Left = node.Left.Delete(maxNode.Value) // 把最大的节点删掉

				node.Left.UpdateHeight() // 删除后要更新该子树高度
			} else {
				minNode := node.Right
				for minNode.Right != nil {
					minNode = minNode.Left
				}
				// 最小值的节点替换被删除节点
				node.Value = minNode.Value
				node.Times = minNode.Times

				node.Left = node.Left.Delete(minNode.Value) // 把最小的节点删掉

				node.Left.UpdateHeight() // 删除后要更新该子树高度
			}
		} else {
			// 只有左子树或只有右子树
			// 只有一个子树，该子树也只是一个节点，将该节点替换被删除的节点，然后置子树为空
			if node.Left != nil {
				//第三种情况，删除的节点只有左子树，因为树的特征，可以知道左子树其实就只有一个节点，它本身，否则高度差就等于2了。
				node.Value = node.Left.Value
				node.Times = node.Left.Times
				node.Height = 1
				node.Left = nil
			} else if node.Right != nil {
				//第四种情况，删除的节点只有右子树，因为树的特征，可以知道右子树其实就只有一个节点，它本身，否则高度差就等于2了。
				node.Value = node.Right.Value
				node.Times = node.Right.Times
				node.Height = 1
				node.Right = nil
			}
		}
		return node
	}

	var newNode *AVLTreeNode
	if node.BalanceFactor() == 2 {
		if node.Left.BalanceFactor() >= 0 {
			newNode = RightRotation(node)
		} else {
			newNode = LeftRightRotation(node)
		}
	} else if node.BalanceFactor() == -2 {
		if node.Right.BalanceFactor() <= 0 {
			newNode = RightRotation(node)
		} else {
			newNode = RightLeftRotation(node)
		}
	}

	if newNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newNode.UpdateHeight()
		return newNode
	}
}
