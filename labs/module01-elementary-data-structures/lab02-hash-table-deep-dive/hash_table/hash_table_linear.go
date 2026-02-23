package hash_table

type LinearHashTable struct {
	tbl []KVPair
	f HashFunc
	m int
}

func (htbl *LinearHashTable) Init(t HashFuncType, m int) {
	htbl.f = SelectHashFunc(t)
	htbl.tbl = make([]KVPair, m)
	htbl.m = m
}

func (htbl *LinearHashTable) Insert(k, v int) {
	idx := htbl.f(k, htbl.m)
	for range htbl.m {
		val := htbl.tbl[idx]
		if !val.filled || (val.key == k && !val.tombstone) {
			break
		}
		idx = (idx + 1) % htbl.m
	}

	htbl.tbl[idx] = KVPair{ key: k, val: v, tombstone: false, filled: true }
}

func (htbl *LinearHashTable) Search(k int) int {
	idx := htbl.f(k, htbl.m)
	var val int
	for range htbl.m {
		v := htbl.tbl[idx]
		if v.filled && v.key == k && !v.tombstone {
			val = v.val
			break
		}
		idx = (idx + 1) % htbl.m
	}

	return val
}

func (htbl *LinearHashTable) Delete(k int) {
	idx := htbl.f(k, htbl.m)
	for range htbl.m {
		v := htbl.tbl[idx]
		if v.filled && v.key == k && !v.tombstone {
			htbl.tbl[idx].tombstone = true
			htbl.tbl[idx].filled = false
			htbl.tbl[idx].key = 0
			htbl.tbl[idx].val = 0
			break
		}
		idx = (idx + 1) % htbl.m
	}
}
