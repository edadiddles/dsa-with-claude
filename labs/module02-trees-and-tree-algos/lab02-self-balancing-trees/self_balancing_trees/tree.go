package self_balancing_tree

type TreeNode struct {
	parent *TreeNode
	val int
	left *TreeNode
	right *TreeNode
}

type ColoredTreeNode struct {
	p *ColoredTreeNode
	val int
	left *ColoredTreeNode
	right *ColoredTreeNode
	color NodeColor
}

type NodeColor int

const (
	BLACK NodeColor = iota
	RED
)
