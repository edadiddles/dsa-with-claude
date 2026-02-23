package hash_table

import (
	"testing"
)

var M = 5

func TestLinearDivisionHash(t *testing.T) {
	dict := LinearHashTable{}
	dict.Init(DIVISION, 10000)
	for k := range M {
		v := k*2
		dict.Insert(k, v)
	}

	for k := range M {
		if dict.Search(k) != 2*k {
			t.Error("insert failed")
		}
	}

	for k := range M {
		dict.Delete(k)
	}

	for k := range M {
		if dict.Search(k) != 0 {
			t.Error("delete failed")
		}
	}
}

func TestLinearMultiplicationHash(t *testing.T) {
	dict := LinearHashTable{}
	dict.Init(MULTIPLICATION, 10000)
	for k := range M {
		v := k*2
		dict.Insert(k, v)
	}

	for k := range M {
		if dict.Search(k) != 2*k {
			t.Error("insert failed", k, dict.Search(k))
		}
	}

	for k := range M {
		dict.Delete(k)
	}

	for k := range M {
		if dict.Search(k) != 0 {
			t.Error("delete failed")
		}
	}
}
