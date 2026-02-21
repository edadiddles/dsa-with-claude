package hash_table

import (
	"testing"
)

func TestBasicDivisionHash(t *testing.T) {
	dict := ChainingHashTable{}
	dict.Init(DIVISION, 10)
	for k := range 1000 {
		v := k*2
		dict.Insert(k, v)
	}

	for k := range 1000 {
		if dict.Search(k) != 2*k {
			t.Error("failed")
		}
	}

	for k := range 1000 {
		dict.Delete(k)
	}

	for k := range 1000 {
		if dict.Search(k) != 0 {
			t.Error("delete failed")
		}
	}
}

func TestBasicMultiplicationHash(t *testing.T) {
	dict := ChainingHashTable{}
	dict.Init(MULTIPLICATION, 10)
	for k := range 1000 {
		v := k*2
		dict.Insert(k, v)
	}

	for k := range 1000 {
		if dict.Search(k) != 2*k {
			t.Error("failed")
		}
	}

	for k := range 1000 {
		dict.Delete(k)
	}

	for k := range 1000 {
		if dict.Search(k) != 0 {
			t.Error("delete failed")
		}
	}
}
