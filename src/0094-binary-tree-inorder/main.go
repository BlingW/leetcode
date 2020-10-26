package main

import "fmt"

// 这里含有多道二叉树的题，包括前中后序遍历搜索及删除
type treeNode struct {
	val   int64
	count int
	left  *treeNode
	right *treeNode
}

type normalBinaryTree struct {
	root *treeNode
	len  int
}

func main() {
	tree := NewBinaryTree()
	tree.Inserts([]int64{50, 60,55,65})
	fmt.Println(tree.InorderTraversal())
	fmt.Println(tree.Search(50))
	fmt.Println(tree.Delete(50))
	fmt.Println(tree.InorderTraversal())
	fmt.Println(tree.Search(50))
}

func NewBinaryTree() *normalBinaryTree {
	return &normalBinaryTree{
		nil, 0,
	}
}

func initTreeNode(i int64) *treeNode {
	return &treeNode{
		i, 1, nil, nil,
	}
}

func (nbt *normalBinaryTree) Insert(i int64) *normalBinaryTree {
	if nbt.root == nil {
		nbt.root = initTreeNode(i)
		return nbt
	}
	nbt.root.insert(i)
	return nbt
}

func (nbt *normalBinaryTree) Inserts(l []int64) *normalBinaryTree {
	if len(l) == 0 {
		return nbt
	}
	if nbt.root == nil {
		nbt.root = initTreeNode(l[0])
	}
	for i := 1; i < len(l); i++ {
		nbt.root.insert(l[i])
	}
	return nbt
}

func (nbt *normalBinaryTree) Search(i int64) int {
	if nbt.root == nil {
		return 0
	}
	return nbt.root.search(i)
}

func (nbt *normalBinaryTree) InorderTraversal() (res []int64) {
	if nbt.root == nil {
		return nil
	}
	nbt.root.inorder(&res)

	return res
}

func (nbt *normalBinaryTree) PreorderTraversal() (res []int64) {
	if nbt.root == nil {
		return nil
	}
	nbt.root.preorder(&res)

	return res
}

func (nbt *normalBinaryTree) PostorderTraversal() (res []int64) {
	if nbt.root == nil {
		return nil
	}
	nbt.root.postorder(&res)

	return res
}

func (tn *treeNode) insert(i int64) {
	switch {
	case i < tn.val:
		if tn.left == nil {
			tn.left = initTreeNode(i)
			return
		}
		tn.left.insert(i)
	case i > tn.val:
		if tn.right == nil {
			tn.right = initTreeNode(i)
			return
		}
		tn.right.insert(i)
	default:
		tn.count++
	}
}

func (tn *treeNode) inorder(res *[]int64) {
	if tn != nil {
		tn.left.inorder(res)
		for i := 1; i <= tn.count; i++ {
			*res = append(*res, tn.val)
		}
		tn.right.inorder(res)
	}
}

func (tn *treeNode) preorder(res *[]int64) {
	if tn != nil {
		for i := 1; i <= tn.count; i++ {
			*res = append(*res, tn.val)
		}
		tn.left.preorder(res)
		tn.right.preorder(res)
	}
}

func (tn *treeNode) postorder(res *[]int64) {
	if tn != nil {
		tn.left.postorder(res)
		tn.right.postorder(res)
		for i := 1; i <= tn.count; i++ {
			*res = append(*res, tn.val)
		}
	}
}

func (tn *treeNode) search(i int64) int {
	cur := tn
	for {
		switch {
		case i < cur.val:
			if cur.left != nil {
				cur = cur.left
				continue
			}
		case i > cur.val:
			if cur.right != nil {
				cur = cur.right
				continue
			}
		default:
			return cur.count
		}
		break
	}
	return 0
}

func (nbt *normalBinaryTree) Delete(i int64) int {
	if nbt.root == nil {
		return 0
	}
	parent, cur := nbt.root, nbt.root
	isLeft := true
	for {
		if i < cur.val {
			if cur.left != nil {
				if cur.left.val == i {
					parent = cur
					cur = cur.left
					break
				}
				cur = cur.left
				continue
			}
		} else if i > cur.val {
			if cur.right != nil {
				if cur.right.val == i {
					parent = cur
					cur = cur.right
					isLeft = false
					break
				}
				cur = cur.right
				continue
			}
		} else {
			break
		}

		return 0
	}
	// 要删除的节点是叶子节点
	if cur.left == nil && cur.right == nil {
		if cur == nbt.root {
			// 特殊情况，只有根节点且根节点就被删除
			nbt.root = nil
		} else if isLeft {
			parent.left = nil
		} else {
			parent.right = nil
		}
		// 要删的节点只有左节点
	} else if cur.right == nil {
		if cur == nbt.root {
			nbt.root = cur.left
		} else if isLeft {
			parent.left = cur.left
		} else {
			parent.right = cur.left
		}
		// 要删的节点只有右节点
	} else if cur.left == nil {
		if cur == nbt.root {
			nbt.root = cur.right
		} else if isLeft {
			parent.left = cur.right
		} else {
			parent.right = cur.right
		}
	} else {
		// 最复杂的情况，既有左节点又有右节点
		suc := cur.getSuccessor()
		if cur == nbt.root {
			nbt.root = suc
		} else if isLeft {
			parent.left = suc
		} else {
			parent.right = suc
		}
		suc.left = cur.left
	}

	return cur.count
}

func (tn *treeNode) getSuccessor() *treeNode {
	var sucParent *treeNode = nil
	suc, cur := tn.right, tn.right
	for cur != nil {
		sucParent = suc
		suc = cur
		cur = cur.left
	}
	if suc != tn.right {
		sucParent.left = suc.right
		suc.right = tn.right
	}
	return suc
}
