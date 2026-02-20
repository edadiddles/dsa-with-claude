package dynamic_array

import (
	"testing"
)

func TestBasicPush(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(10, 1.0, ADD)
	arr.Push(5)

	if arr.Size() != 1 {
		t.Error("push length not correct")
	}
	if arr.Get(0) != 5 {
		t.Error("push value not correct")
	}
}

func TestBasicPop(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(10, 1.0, ADD)
	arr.Push(5)
	arr.Push(6)

	if arr.Size() != 2 {
		t.Error("length not correct")
	}
	if arr.Pop() != 6 {
		t.Error("pop value not correct")
	}
	if arr.Size() != 1 {
		t.Error("pop did not reduce length")
	}
}

func TestBasicGet(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(10, 1.0, ADD)
	arr.Push(5)
	arr.Push(6)

	if arr.Size() != 2 {
		t.Error("length not correct")
	}
	if arr.Get(0) != 5 {
		t.Error("get value 0 not correct")
	}
	if arr.Get(1) != 6 {
		t.Error("get value 1 not correct")
	}
}

func TestBasicSet(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(10, 1.0, ADD)
	arr.Push(5)
	arr.Push(6)
	arr.Set(0, 1)
	arr.Set(1, 2)

	if arr.Size() != 2 {
		t.Error("length not correct")
	}
	if arr.Get(0) != 1 {
		t.Error("get value 0 not correct")
	}
	if arr.Get(1) != 2 {
		t.Error("get value 1 not correct")
	}
}

func TestBasicInsert(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(10, 1.0, ADD)
	arr.Push(5)
	arr.Push(6)
	arr.Insert(1, 1)
	arr.Insert(0, 2)

	if arr.Size() != 4 {
		t.Error("length not correct")
	}
	if arr.Get(0) != 2 {
		t.Error("get value 0 not correct")
	}
	if arr.Get(1) != 5 {
		t.Error("get value 1 not correct")
	}
	if arr.Get(2) != 1 {
		t.Error("get value 2 not correct")
	}
	if arr.Get(3) != 6 {
		t.Error("get value 3 not correct")
	}
}

func TestBasicDelete(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(10, 1.0, ADD)
	arr.Push(5)
	arr.Push(6)
	arr.Delete(0)

	if arr.Size() != 1 {
		t.Error("length not correct")
	}
	if arr.Get(0) != 6 {
		t.Error("get value 0 not correct")
	}
}

func TestBasicSize(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(10, 1.0, ADD)
	arr.Push(5)
	arr.Push(6)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)

	if arr.Size() != 5 {
		t.Error("length not correct")
	}
}

func TestBasicCapacity(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(10, 1.0, ADD)
	arr.Push(5)
	arr.Push(6)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)

	if arr.Capacity() != 10 {
		t.Error("capacity not correct")
	}
}

func TestDynamicGrowthMult2(t *testing.T) {	
	arr := DynamicArray{}

	arr.Init(2, 2.0, MULTIPLY)
	if arr.Capacity() != 2 {
		t.Error("capacity not correct")
	}

	arr.Push(5)
	arr.Push(6)
	arr.Push(7)	
	if arr.Capacity() != 4 {
		t.Error("capacity not correct")
	}

	arr.Push(7)
	arr.Push(7)
	if arr.Capacity() != 8 {
		t.Error("capacity not correct")
	}
}

func TestDynamicGrowthMult3(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(2, 3.0, MULTIPLY)
	if arr.Capacity() != 2 {
		t.Error("capacity not correct")
	}

	arr.Push(5)
	arr.Push(5)
	arr.Push(6)
	arr.Push(7)	
	if arr.Capacity() != 6 {
		t.Error("capacity not correct")
	}

	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	if arr.Capacity() != 18 {
		t.Error("capacity not correct")
	}


	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	arr.Push(7)
	if arr.Capacity() != 54 {
		t.Error("capacity not correct")
	}
}

func TestDyanmicGrowthMult1_5(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(3, 1.5, MULTIPLY)
	if arr.Capacity() != 3 {
		t.Error("capacity not correct")
	}

	arr.Push(5)
	arr.Push(5)
	arr.Push(6)
	arr.Push(7)	
	if arr.Capacity() != 4 {
		t.Error("capacity not correct")
	}

	arr.Push(7)
	if arr.Capacity() != 6 {
		t.Error("capacity not correct")
	}


	arr.Push(7)
	arr.Push(7)
	if arr.Capacity() != 9 {
		t.Error("capacity not correct")
	}
}

func TestDynamicGrowthAdd5(t *testing.T) {
	arr := DynamicArray{}

	arr.Init(3, 5, ADD)
	if arr.Capacity() != 3 {
		t.Error("capacity not correct")
	}

	arr.Push(5)
	arr.Push(5)
	arr.Push(6)
	arr.Push(7)	
	if arr.Capacity() != 8 {
		t.Error("capacity not correct")
	}

	arr.Push(7)
	arr.Push(5)
	arr.Push(5)
	arr.Push(5)
	arr.Push(5)
	if arr.Capacity() != 13 {
		t.Error("capacity not correct")
	}
}
