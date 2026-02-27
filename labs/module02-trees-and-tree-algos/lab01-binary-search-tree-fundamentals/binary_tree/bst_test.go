package binary_tree

import (
	"testing"
)

func TestBSTBasic(t *testing.T) {
	bst := BinarySearchTree{}

	bst.Insert(15)

	if bst.Size() != 1 {
		t.Error(bst.Size())
		t.Error("size is wrong")
	}
	if bst.Height() != 1 {
		t.Error(bst.Height())
		t.Error("height is wrong")
	}

	bst.Insert(6)
	bst.Insert(18)
	if bst.Size() != 3 {
		t.Error(bst.Size())
		t.Error("size is wrong")
	}
	if bst.Height() != 2 {
		t.Error(bst.Height())
		t.Error("height is wrong")
	}
	
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(17)
	bst.Insert(20)
	if bst.Size() != 7 {
		t.Error(bst.Size())
		t.Error("size is wrong")
	}
	if bst.Height() != 3 {
		t.Error(bst.Height())
		t.Error("height is wrong")
	}

	bst.Delete(9)
	if bst.Size() != 7 {
		t.Error(bst.Size())
		t.Error("size is wrong")
	}
	if bst.Height() != 3 {
		t.Error(bst.Height())
		t.Error("height is wrong")
	}

	bst.Delete(6)
	if bst.Size() != 6 {
		t.Error(bst.Size())
		t.Error("size is wrong")
	}
	if bst.Height() != 3 {
		t.Error(bst.Height())
		t.Error("height is wrong")
	}

	bst.Delete(15)
	if bst.Size() != 5 {
		t.Error(bst.Size())
		t.Error("size is wrong")
	}
	if bst.Height() != 3 {
		t.Error(bst.Height())
		t.Error("height is wrong")
	}

	bst.Delete(3)
	bst.Delete(7)
	bst.Delete(17)
	bst.Delete(20)
	if bst.Size() != 1 {
		t.Error(bst.Size())
		t.Error("size is wrong")
	}
	if bst.Height() != 1 {
		t.Error(bst.Height())
		t.Error("height is wrong")
	}

}
