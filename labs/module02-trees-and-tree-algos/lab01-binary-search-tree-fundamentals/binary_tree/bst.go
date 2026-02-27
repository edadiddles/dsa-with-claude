package binary_tree

type BinarySearchTree struct {
	tree *TreeNode
}

type TreeNode struct {
	p *TreeNode
	val int
	left *TreeNode
	right *TreeNode
}

func (bst *BinarySearchTree) Insert(val int) {
	bst.tree = bst.r_insert(bst.tree, val)

}

func (bst *BinarySearchTree) Search(val int) *TreeNode {
	return bst.r_search(bst.tree, val)
}

func (bst *BinarySearchTree) Delete(val int) {
	n := bst.Search(val)
	if n == nil {
		return
	}

	if n.left == nil {
		bst.transplant(n, n.right)
	} else if n.right == nil {
		bst.transplant(n, n.left)
	} else {
		u := bst.r_minimum(n.right)
		if u != n.right {
			bst.transplant(u, u.right)
			u.right = n.right
			u.right.p = u
		}
		bst.transplant(n, u)
		u.left = n.left
		u.left.p = u
	}
}

func (bst *BinarySearchTree) Minimum() *TreeNode {
	return bst.r_minimum(bst.tree)
}

func (bst *BinarySearchTree) Maximum() *TreeNode {
	return bst.r_maximum(bst.tree)
}

func (bst *BinarySearchTree) Height() int {
	return bst.r_height(bst.tree, 0, 0)
}

func (bst *BinarySearchTree) Size() int {
	return bst.r_size(bst.tree, 0)
}

func (bst *BinarySearchTree) r_insert(n *TreeNode, val int) *TreeNode {
	if n == nil {
		return &TreeNode{val: val}
	}

	if val < n.val {
		n.left = bst.r_insert(n.left, val)
		n.left.p = n
	} else {
		n.right = bst.r_insert(n.right, val)
		n.right.p = n
	}

	return n

}

func (bst *BinarySearchTree) r_search(n *TreeNode, val int) *TreeNode {
	if n == nil || n.val == val {
		return n
	}

	if val < n.val {
		return bst.r_search(n.left, val)
	} else {
		return bst.r_search(n.right, val)
	}
}

func (bst *BinarySearchTree) r_minimum(n *TreeNode) *TreeNode {
	if n.left == nil {
		return n
	}

	return n.left
}

func (bst *BinarySearchTree) r_maximum(n *TreeNode) *TreeNode {
	if n.right == nil {
		return n
	}

	return n.right
}

func (bst *BinarySearchTree) r_height(n *TreeNode, curr_height int, max_height int) int {
	if n == nil {
		return max_height
	}

	curr_height += 1
	if curr_height > max_height {
		max_height = curr_height
	}
	max_height_left := bst.r_height(n.left, curr_height, max_height)
	max_height_right := bst.r_height(n.right, curr_height, max_height)

	if max_height_left > max_height_right {
		return max_height_left
	} else {
		return max_height_right
	}
}

func (bst *BinarySearchTree) r_size(n *TreeNode, s int) int {
	if n == nil {
		return s
	}
	s+=1
	s += bst.r_size(n.left, 0)
	s += bst.r_size(n.right, 0)

	return s
}

func (bst *BinarySearchTree) transplant(u, v *TreeNode) {
	if u.p == nil {
		bst.tree = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}

	if v != nil {
		v.p = u.p
	}
}
