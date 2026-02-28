package self_balancing_tree

var NIL = &ColoredTreeNode{ color: BLACK }

type RBTree struct {
	root *ColoredTreeNode
}

func (T *RBTree) Insert(val int) {
	z := &ColoredTreeNode{ p: NIL, val: val, color: BLACK, left: NIL, right: NIL}
	if T.root == nil {
		T.root = z
		return
	}

	x := T.root
	y := NIL
	for x != NIL {
		y = x
		if z.val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.p = y
	if y == NIL {
		T.root = z
	} else if z.val < y.val {
		y.left = z
	} else {
		y.right = z
	}

	z.color = RED
	T.insert_fixup(z)

}

func (T *RBTree) Search(val int) *ColoredTreeNode {
	if T.root == nil {
		return NIL
	}

	x := T.root
	for x != NIL {
		if x.val == val {
			break
		} else if val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}

	return x
}

func (T *RBTree) Delete(val int) {
	z := T.Search(val)
	y := z
	y_orig_color := y.color

	var x *ColoredTreeNode
	if z.left == NIL {
		x = z.right
		T.transplant(z, z.right)
	} else if z.right == NIL {
		x = z.left
		T.transplant(z, z.left)
	} else {
		y = T.minimum(z.right)
		y_orig_color = y.color
		x = y.right
		if y != z.right {
			T.transplant(y, y.right)
			y.right = z.right
			y.right.p = y
		} else {
			x.p = y
		}
		T.transplant(z, y)
		y.left = z.left
		y.left.p = y
		y.color = z.color
	}

	if y_orig_color == BLACK {
		T.delete_fixup(x)
	}
}

func (T *RBTree) minimum(x *ColoredTreeNode) *ColoredTreeNode {
	if x == nil {
		return NIL
	}

	for x.left != NIL {
		x = x.left
	}

	return x
}

func (T *RBTree) maximum(x *ColoredTreeNode) *ColoredTreeNode {
	if x == nil {
		return NIL
	}

	for x.right != NIL {
		x = x.right
	}

	return x
}

func (T *RBTree) left_rotate(x *ColoredTreeNode) {
	y := x.right
	x.right = y.left
	if y.left != NIL {
		y.left.p = x
	}
	y.p = x.p
	if x.p == NIL {
		T.root = y
	} else if x == x.p.left {
		x.p.left = y
	} else {
		x.p.right = y
	}
	y.left = x
	x.p = y
}

func (T *RBTree) right_rotate(y *ColoredTreeNode) {
	x := y.left
	y.left = x.right
	if x.right != NIL {
		x.right.p = y
	}
	x.p = y.p
	if y.p == NIL {
		T.root = x
	} else if y == y.p.left {
		y.p.left = x
	} else {
		y.p.right = x
	}
	x.right = y
	y.p = x
}

func (T *RBTree) transplant(u, v *ColoredTreeNode) {
	if u.p == NIL {
		T.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	v.p = u.p
}

func (T *RBTree) insert_fixup(z *ColoredTreeNode) {
	for z.p.color == RED {
		if z.p == z.p.p.left {
			y := z.p.p.right
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.right {
					z = z.p
					T.left_rotate(z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				T.right_rotate(z.p.p)
			}
		} else {
			y := z.p.p.left
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.left {
					z = z.p
					T.right_rotate(z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				T.left_rotate(z.p.p)
			}
		}
	}
	T.root.color = BLACK
}

func (T *RBTree) delete_fixup(x *ColoredTreeNode) {
	for x != T.root && x.color == BLACK {
		if x == x.p.left {
			w := x.p.right
			if w.color == RED {
				w.color = BLACK
				x.p.color = RED
				T.left_rotate(x.p)
				w = x.p.right
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.p
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					T.right_rotate(w)
					w = x.p.right
				}
				w.color = x.p.color
				x.p.color = BLACK
				w.right.color = BLACK
				T.left_rotate(x.p)
				x = T.root
			}
		} else {
			w := x.p.left
			if w.color == RED {
				w.color = BLACK
				x.p.color = RED
				T.right_rotate(x.p)
				w = x.p.left
			}
			if w.right.color == BLACK && w.left.color == BLACK {
				w.color = RED
				x = x.p
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					T.left_rotate(w)
					w = x.p.left
				}
				w.color = x.p.color
				x.p.color = BLACK
				w.left.color = BLACK
				T.right_rotate(x.p)
				x = T.root
			}
		}	
	}
	x.color = BLACK
}
